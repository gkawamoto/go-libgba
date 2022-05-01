package main

// Inspired by https://github.com/pret/pokefirered/blob/master/src/isagbprn.c

import (
	"runtime/volatile"
	"unsafe"
)

func main() {
	AGBPrintInit()
	AGBPrint("hello world!")

	select {}
}

/*
#define AGB_PRINT_FLUSH_ADDR 0x9FE209D
#define AGB_PRINT_STRUCT_ADDR 0x9FE20F8
#define AGB_PRINT_PROTECT_ADDR 0x9FE2FFE
#define WSCNT_DATA (WAITCNT_PHI_OUT_16MHZ | WAITCNT_WS0_S_2 | WAITCNT_WS0_N_4)

static void AGBPutcInternal(const char cChr)
{
    volatile struct AGBPrintStruct *pPrint = (struct AGBPrintStruct *)AGB_PRINT_STRUCT_ADDR;
    u16 *pPrintBuf = (u16 *)(0x8000000 + (pPrint->m_nBank << 16));
    u16 *pProtect = (u16 *)AGB_PRINT_PROTECT_ADDR;
    u16 nData = pPrintBuf[pPrint->m_nPut / 2];
    *pProtect = 0x20;
    nData = (pPrint->m_nPut & 1) ? (nData & 0xFF) | (cChr << 8) : (nData & 0xFF00) | cChr;
    pPrintBuf[pPrint->m_nPut / 2] = nData;
    pPrint->m_nPut++;
    *pProtect = 0;
}

void AGBPutc(const char cChr)
{
    u16 *pWSCNT = (u16 *)REG_ADDR_WAITCNT;
    u16 nOldWSCNT = *pWSCNT;
    struct AGBPrintStruct *pPrint;
    *pWSCNT = WSCNT_DATA;
    AGBPutcInternal(cChr);
    *pWSCNT = nOldWSCNT;
    pPrint = (struct AGBPrintStruct *)AGB_PRINT_STRUCT_ADDR;
    if (pPrint->m_nPut == ((pPrint->m_nGet - 1) & 0xFFFF))
        AGBPrintFlush1Block();
}
*/

const (
	WAITCNT_PHI_OUT_16MHZ = (3 << 11)
	WAITCNT_WS0_S_2       = (0 << 4)
	WAITCNT_WS0_N_4       = (0 << 2)
	WSCNT_DATA            = (WAITCNT_PHI_OUT_16MHZ | WAITCNT_WS0_S_2 | WAITCNT_WS0_N_4)
)

/*
   u16 m_nRequest;
   u16 m_nBank;
   u16 m_nGet;
   u16 m_nPut;
*/
var AGB_PRINT_STRUCT_ADDR = (*[4]volatile.Register16)(unsafe.Pointer(uintptr(0x9FE20F8)))
var AGB_PRINT_PROTECT_ADDR = (*volatile.Register16)(unsafe.Pointer(uintptr(0x9FE2FFE)))
var REG_ADDR_WAITCNT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000204)))

var REG_ADDR_IME = (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000208)))

/*
void AGBPrintInit(void)
{
    struct AGBPrintStruct *pPrint = (struct AGBPrintStruct *)AGB_PRINT_STRUCT_ADDR;
    u16 *pWSCNT = (u16 *)REG_ADDR_WAITCNT;
    u16 *pProtect = (u16 *)AGB_PRINT_PROTECT_ADDR;
    u16 nOldWSCNT = *pWSCNT;
    *pWSCNT = WSCNT_DATA;
    *pProtect = 0x20;
    pPrint->m_nRequest = pPrint->m_nGet = pPrint->m_nPut = 0;
    pPrint->m_nBank = 0xFD;
    *pProtect = 0;
    *pWSCNT = nOldWSCNT;
}
*/
func AGBPrintInit() {
	nOldWSCNT := REG_ADDR_WAITCNT.Get()
	REG_ADDR_WAITCNT.Set(WSCNT_DATA)
	AGB_PRINT_PROTECT_ADDR.Set(0x20)
	AGB_PRINT_STRUCT_ADDR[0].Set(0)
	AGB_PRINT_STRUCT_ADDR[1].Set(0xFD)
	AGB_PRINT_STRUCT_ADDR[2].Set(0)
	AGB_PRINT_STRUCT_ADDR[3].Set(0)
	AGB_PRINT_PROTECT_ADDR.Set(0)
	REG_ADDR_WAITCNT.Set(nOldWSCNT)
}

/*
void AGBPrint(const char *pBuf)
{
    struct AGBPrintStruct *pPrint = (struct AGBPrintStruct *)AGB_PRINT_STRUCT_ADDR;
    u16 *pWSCNT = (u16 *)REG_ADDR_WAITCNT;
    u16 nOldWSCNT = *pWSCNT;
    *pWSCNT = WSCNT_DATA;
    while (*pBuf)
    {
        AGBPutc(*pBuf);
        pBuf++;
    }
    *pWSCNT = nOldWSCNT;
}
*/

func AGBPrint(value string) {
	nOldWSCNT := REG_ADDR_WAITCNT.Get()
	REG_ADDR_WAITCNT.Set(WSCNT_DATA)
	for _, c := range value {
		AGBPutc(c)
	}
	REG_ADDR_WAITCNT.Set(nOldWSCNT)
}

/*
void AGBPutc(const char cChr)
{
    u16 *pWSCNT = (u16 *)REG_ADDR_WAITCNT;
    u16 nOldWSCNT = *pWSCNT;
    struct AGBPrintStruct *pPrint;
    *pWSCNT = WSCNT_DATA;
    AGBPutcInternal(cChr);
    *pWSCNT = nOldWSCNT;
    pPrint = (struct AGBPrintStruct *)AGB_PRINT_STRUCT_ADDR;
    if (pPrint->m_nPut == ((pPrint->m_nGet - 1) & 0xFFFF))
        AGBPrintFlush1Block();
}
*/
func AGBPutc(value rune) {
	// u16 nOldWSCNT = *pWSCNT;
	nOldWSCNT := REG_ADDR_WAITCNT.Get()

	// *pWSCNT = WSCNT_DATA;
	REG_ADDR_WAITCNT.Set(WSCNT_DATA)

	// AGBPutcInternal(cChr);
	AGBPutcInternal(value)

	// *pWSCNT = nOldWSCNT;
	REG_ADDR_WAITCNT.Set(nOldWSCNT)

	/*
	   pPrint = (struct AGBPrintStruct *)AGB_PRINT_STRUCT_ADDR;
	   if (pPrint->m_nPut == ((pPrint->m_nGet - 1) & 0xFFFF))
	       AGBPrintFlush1Block();
	*/
	if AGB_PRINT_STRUCT_ADDR[3].Get() == (AGB_PRINT_STRUCT_ADDR[2].Get()-1)&0xFFFF {
		AGBPrintFlush1Block()
	}
}

/*
void AGBPrintFlush1Block(void)
{
    AGBPrintTransferDataInternal(FALSE);
}
*/
func AGBPrintFlush1Block() {
	AGBPrintTransferDataInternal(false)
}

/*
static void AGBPrintTransferDataInternal(u32 bAllData)
{
    LPFN_PRINT_FLUSH lpfnFuncFlush;
    u16 *pIME;
    u16 nIME;
    u16 *pWSCNT;
    u16 nOldWSCNT;
    u16 *pProtect;
    volatile struct AGBPrintStruct *pPrint;

    pProtect = (u16 *)AGB_PRINT_PROTECT_ADDR;
    pPrint = (struct AGBPrintStruct *)AGB_PRINT_STRUCT_ADDR;
    lpfnFuncFlush = (LPFN_PRINT_FLUSH)AGB_PRINT_FLUSH_ADDR;
    pIME = (u16 *)REG_ADDR_IME;
    nIME = *pIME;
    pWSCNT = (u16 *)REG_ADDR_WAITCNT;
    nOldWSCNT = *pWSCNT;
    *pIME = nIME & ~1;
    *pWSCNT = WSCNT_DATA;

    if (bAllData)
    {
        while (pPrint->m_nPut != pPrint->m_nGet)
        {
            *pProtect = 0x20;
            lpfnFuncFlush();
            *pProtect = 0;
        }
    }
    else if (pPrint->m_nPut != pPrint->m_nGet)
    {
        *pProtect = 0x20;
        lpfnFuncFlush();
        *pProtect = 0;
    }

    *pWSCNT = nOldWSCNT;
    *pIME = nIME;
}
*/
func AGBPrintTransferDataInternal(bAllData bool) {
	lpfnFuncFlush := (*func())(unsafe.Pointer(uintptr(0x9FE209D)))
	// nIME = *pIME;
	nIME := REG_ADDR_IME.Get()
	nOldWSCNT := REG_ADDR_WAITCNT.Get()

	// *pIME = nIME & ~1;
	REG_ADDR_IME.Set(nIME &^ 1)
	REG_ADDR_WAITCNT.Set(WSCNT_DATA)

	if bAllData {
		for AGB_PRINT_STRUCT_ADDR[3].Get() != AGB_PRINT_STRUCT_ADDR[2].Get() {
			AGB_PRINT_PROTECT_ADDR.Set(0x20)
			(*lpfnFuncFlush)()
			AGB_PRINT_PROTECT_ADDR.Set(0)
		}
	} else if AGB_PRINT_STRUCT_ADDR[3].Get() != AGB_PRINT_STRUCT_ADDR[2].Get() {
		AGB_PRINT_PROTECT_ADDR.Set(0x20)
		(*lpfnFuncFlush)()
		AGB_PRINT_PROTECT_ADDR.Set(0)
	}

	REG_ADDR_WAITCNT.Set(nOldWSCNT)
	REG_ADDR_IME.Set(nIME)
}

/*
static void AGBPutcInternal(const char cChr)
{
    volatile struct AGBPrintStruct *pPrint = (struct AGBPrintStruct *)AGB_PRINT_STRUCT_ADDR;
    u16 *pPrintBuf = (u16 *)(0x8000000 + (pPrint->m_nBank << 16));
    u16 *pProtect = (u16 *)AGB_PRINT_PROTECT_ADDR;
    u16 nData = pPrintBuf[pPrint->m_nPut / 2];
    *pProtect = 0x20;
    nData = (pPrint->m_nPut & 1) ? (nData & 0xFF) | (cChr << 8) : (nData & 0xFF00) | cChr;
    pPrintBuf[pPrint->m_nPut / 2] = nData;
    pPrint->m_nPut++;
    *pProtect = 0;
}
*/
func AGBPutcInternal(value rune) {

	nPut := AGB_PRINT_STRUCT_ADDR[3].Get()

	var pPrintBuf = (*volatile.Register16)(unsafe.Pointer(uintptr(0x8000000 + uint(AGB_PRINT_STRUCT_ADDR[1].Get())<<16 + uint(nPut)/2)))
	AGB_PRINT_PROTECT_ADDR.Set(0x20)
	nData := pPrintBuf.Get()

	var newValue uint16
	if nPut&1 == 1 {
		newValue = (nData & 0xFF) | uint16(value<<8)
	} else {
		newValue = (nData & 0xFF00) | uint16(value)
	}
	pPrintBuf.Set(newValue)
	AGB_PRINT_STRUCT_ADDR[3].Set(nPut + 1)
	AGB_PRINT_PROTECT_ADDR.Set(0)
}
