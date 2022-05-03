package agbprint

// #include <stdio.h>
//
// typedef void (*LPFN_PRINT_FLUSH)(void);
//
// // We need this because Tinygo can't handle direct volatile function calls
// void lpfn_func_flush() {
//     LPFN_PRINT_FLUSH lpfnFuncFlush;
//     lpfnFuncFlush = (LPFN_PRINT_FLUSH)0x9FE209D;
//     lpfnFuncFlush();
// }
import "C"

import (
	"fmt"
	"log"
	"runtime/volatile"
	"unsafe"
)

// This file was inspired by https://github.com/pret/pokefirered/blob/master/src/isagbprn.c

// Constants we'll need to set in order to reset state for some register controls in logging facilities
const (
	_WAITCNT_PHI_OUT_16MHZ = (3 << 11)
	_WAITCNT_WS0_S_2       = (0 << 4)
	_WAITCNT_WS0_N_4       = (0 << 2)
	_WSCNT_DATA            = (_WAITCNT_PHI_OUT_16MHZ | _WAITCNT_WS0_S_2 | _WAITCNT_WS0_N_4)
)

var (
	// Struct that contains controls to handle logging facilities
	_AGB_PRINT_STRUCT_ADDR = (*[4]volatile.Register16)(unsafe.Pointer(uintptr(0x9FE20F8)))
	/*
	   [0] = u16 m_nRequest;
	   [1] = u16 m_nBank;
	   [2] = u16 m_nGet;
	   [3] = u16 m_nPut;
	*/

	// Logging write protection
	_AGB_PRINT_PROTECT_ADDR = (*volatile.Register16)(unsafe.Pointer(uintptr(0x9FE2FFE)))

	// Waiting register controller
	_REG_ADDR_WAITCNT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000204)))

	_REG_ADDR_IME = (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000208)))
)

type writer struct {
}

func (w *writer) Write(p []byte) (n int, err error) {
	text := string(p)
	Println(text[0 : len(text)-1]) // log package always includes a newline, but that's not necessary in GBA
	agbPrintFlush()                // Flush right away so we see results in the emulator
	return len(p), nil
}

// Init initializes the logging facilities in the emulator and overrides the `log` package output to send it to GBA logging instead
func Init() {
	agbPrintInit() // Initializing controller registers

	log.SetFlags(log.Lmsgprefix)
	log.SetOutput(&writer{})
}

// Println outputs the text to the emulator logs
func Println(text string) {
	agbPrint(text)
}

// Printf formats the message using `fmt` and calls Println to do the rest
func Printf(format string, a ...interface{}) {
	Println(fmt.Sprintf(format, a...))
}

// agbPrintInit initializes the controller register values
func agbPrintInit() {
	nOldWSCNT := _REG_ADDR_WAITCNT.Get()
	_REG_ADDR_WAITCNT.Set(_WSCNT_DATA)
	_AGB_PRINT_PROTECT_ADDR.Set(0x20)
	_AGB_PRINT_STRUCT_ADDR[0].Set(0)
	_AGB_PRINT_STRUCT_ADDR[1].Set(0xFD)
	_AGB_PRINT_STRUCT_ADDR[2].Set(0)
	_AGB_PRINT_STRUCT_ADDR[3].Set(0)
	_AGB_PRINT_PROTECT_ADDR.Set(0)
	_REG_ADDR_WAITCNT.Set(nOldWSCNT)
}

// agbPrint loops through a string's characters and sends it byte by byte to the logging memory bank
func agbPrint(value string) {
	nOldWSCNT := _REG_ADDR_WAITCNT.Get()
	_REG_ADDR_WAITCNT.Set(_WSCNT_DATA)
	for _, c := range value {
		agbPutc(byte(c))
	}
	_REG_ADDR_WAITCNT.Set(nOldWSCNT)
}

// agbPutc writes and flushes the logging bank in case it moved
func agbPutc(value byte) {
	nOldWSCNT := _REG_ADDR_WAITCNT.Get()
	_REG_ADDR_WAITCNT.Set(_WSCNT_DATA)

	agbPutcInternal(value)

	_REG_ADDR_WAITCNT.Set(nOldWSCNT)

	if _AGB_PRINT_STRUCT_ADDR[3].Get() == (_AGB_PRINT_STRUCT_ADDR[2].Get()-1)&0xFFFF {
		agbPrintFlush1Block()
	}
}

// agbPrintFlush1Block flushes a single block in memory bank
func agbPrintFlush1Block() {
	agbPrintTransferDataInternal(false)
}

// agbPrintTransferDataInternal flushes a single block or the entire bank if bAllData is true
func agbPrintTransferDataInternal(bAllData bool) {
	nIME := _REG_ADDR_IME.Get()
	nOldWSCNT := _REG_ADDR_WAITCNT.Get()

	_REG_ADDR_IME.Set(nIME &^ 1)
	_REG_ADDR_WAITCNT.Set(_WSCNT_DATA)

	switch {
	case bAllData:
		for _AGB_PRINT_STRUCT_ADDR[3].Get() != _AGB_PRINT_STRUCT_ADDR[2].Get() {
			_AGB_PRINT_PROTECT_ADDR.Set(0x20)
			C.lpfn_func_flush()
			_AGB_PRINT_PROTECT_ADDR.Set(0)
		}

	case _AGB_PRINT_STRUCT_ADDR[3].Get() != _AGB_PRINT_STRUCT_ADDR[2].Get():
		_AGB_PRINT_PROTECT_ADDR.Set(0x20)
		C.lpfn_func_flush()
		_AGB_PRINT_PROTECT_ADDR.Set(0)
	}

	_REG_ADDR_WAITCNT.Set(nOldWSCNT)
	_REG_ADDR_IME.Set(nIME)
}

// agbPrintFlush flushes the entire memory bank to the emulator's logging facilities
func agbPrintFlush() {
	agbPrintTransferDataInternal(true)
}

// agbPutcInternal writes the byte to the next position in the current bank
func agbPutcInternal(value byte) {

	nPut := _AGB_PRINT_STRUCT_ADDR[3].Get()

	var pPrintBuf = (*[65536]volatile.Register16)(unsafe.Pointer(uintptr(0x8000000 + uint(_AGB_PRINT_STRUCT_ADDR[1].Get())<<16)))
	_AGB_PRINT_PROTECT_ADDR.Set(0x20)
	nData := pPrintBuf[uint(nPut)/2].Get()

	// We cannot write single bytes in GBA, so in order to write it correctly we need to do some bitwise magic to keep the old value.
	// Maybe we can write two bytes at once instead, optimizing performance and getting rid of (most) bitwise magic
	var newValue uint16
	if nPut&1 == 1 {
		newValue = (nData & 0xFF) | uint16(value)<<8
	} else {
		newValue = (nData & 0xFF00) | uint16(value)
	}
	pPrintBuf[uint(nPut)/2].Set(newValue)
	_AGB_PRINT_STRUCT_ADDR[3].Set(nPut + 1)
	_AGB_PRINT_PROTECT_ADDR.Set(0)
}
