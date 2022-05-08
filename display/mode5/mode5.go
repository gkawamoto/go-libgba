package mode5

import (
	"runtime/volatile"
	"unsafe"
)

// Video memory (0600:0000h)
var _MEM_VRAM_PAGE1 = (*[20480]volatile.Register16)(unsafe.Pointer(uintptr(0x06000000)))

// Video memory (0600:A000h)
var _MEM_VRAM_PAGE2 = (*[20480]volatile.Register16)(unsafe.Pointer(uintptr(0x0600A000)))

// SetPixel sets a 15bit color in a x/y coordinate
func SetPixel(x, y, color uint16, frontPage bool) {
	if frontPage {
		_MEM_VRAM_PAGE1[x+y*160].Set(color)
	} else {
		_MEM_VRAM_PAGE2[x+y*160].Set(color)
	}
}
