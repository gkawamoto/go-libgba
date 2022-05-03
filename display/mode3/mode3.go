package mode3

import (
	"runtime/volatile"
	"unsafe"
)

// Video memory (0600:0000h)
var _MEM_VRAM = (*[38400]volatile.Register16)(unsafe.Pointer(uintptr(0x06000000)))

// SetPixel sets a 15bit color in a x/y coordinate
func SetPixel(x, y, color uint16) {
	_MEM_VRAM[x+y*240].Set(color)
}
