package display

import (
	"machine"
	"runtime/interrupt"
	"runtime/volatile"
	"sync"
	"unsafe"
)

type DisplayOption uint16

var (
	Mode0 DisplayOption = 0
	Mode1 DisplayOption = 1
	Mode2 DisplayOption = 2
	Mode3 DisplayOption = 3
	Mode4 DisplayOption = 4
	Mode5 DisplayOption = 5

	BG0 DisplayOption = 1 << 8
	BG1 DisplayOption = 1 << 9
	BG2 DisplayOption = 1 << 10
	BG3 DisplayOption = 1 << 11
	Obj DisplayOption = 1 << 12

	BackPage DisplayOption = 1 << 4
)

var vBlankUpdateRefs = []func(){}

var (
	// Blank interrupt request. If set, an interrupt will be fired at VBlank.
	_DSTAT_VBL_IRQ uint16 = 1 << 3
	_DSTAT_HBL_IRQ uint16 = 1 << 4
)

var once = &sync.Once{}

var (
	// Display status (0400:0004h)
	_REG_DISPSTAT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000004)))

	_REG_DISPCNT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000000)))
)

func EnableVBlankInterrupts() {
	_REG_DISPSTAT.SetBits(_DSTAT_VBL_IRQ)
	interrupt.New(machine.IRQ_VBLANK, vBlankUpdate).Enable()
}

func vBlankUpdate(interrupt.Interrupt) {
	for _, ref := range vBlankUpdateRefs {
		ref()
	}
}

func RegisterVBlankUpdate(ref func()) {
	vBlankUpdateRefs = append(vBlankUpdateRefs, ref)
}

func SetDisplayOptions(options ...DisplayOption) {
	var value DisplayOption = 0
	for _, option := range options {
		value = value | option
	}
	_REG_DISPCNT.Set(uint16(value))
}

func ShowPage(frontPage bool) {
	value := _REG_DISPSTAT.Get()
	if frontPage {
		var mask uint16 = ^(uint16(1) << 4)
		_REG_DISPCNT.Set(value & mask)
	} else {
		_REG_DISPSTAT.Set(value | uint16(1)<<4)
	}
}