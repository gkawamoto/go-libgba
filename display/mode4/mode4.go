package mode4

import (
	"runtime/volatile"
	"unsafe"
)

// Video memory page 1 as uint16 registers, as memory cannot be set as uint8 in GBA vram (0600:0000h)
var _MEM_VRAM_PAGE1 = (*[160][120]volatile.Register16)(unsafe.Pointer(uintptr(0x06000000)))

// Video memory page 2 as uint16 registers, as memory cannot be set as uint8 in GBA vram (0600:A000h)
var _MEM_VRAM_PAGE2 = (*[160][120]volatile.Register16)(unsafe.Pointer(uintptr(0x0600A000)))

// Background palette memory (0500:0000h)
var _MEM_BGPALETTE = (*[256]volatile.Register16)(unsafe.Pointer(uintptr(0x05000000)))

// SetPalette defines a 15bit color to the set index. Use color.PackRGB to generate 15bit colors.
func SetPalette(index uint8, color uint16) {
	_MEM_BGPALETTE[index].Set(color)
}

// SetPixel sets a palette color on a point in a x/y coordinate, either in the front or the back page.
func SetPixel(x, y uint8, index uint8, frontPage bool) {
	var mem *[160][120]volatile.Register16
	if frontPage {
		mem = _MEM_VRAM_PAGE1
	} else {
		mem = _MEM_VRAM_PAGE2
	}

	// We get the current value (divided by two because registers are only writable in 2 byte chunks)
	pos := mem[y][x/2]

	// We need the value set before so we don't override it when writing a single byte using 2 bytes.
	dst := pos.Get()

	// Little magic to write either the odd or even byte keeping the other byte value in this 2 byte operation.
	if x&1 == 1 {
		pos.Set(dst&0xFF | uint16(index)<<8)
	} else {
		pos.Set(dst&^0xFF | uint16(index))
	}
}
