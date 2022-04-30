package keys

import (
	"runtime/volatile"
	"unsafe"
)

var _REG_KEYCNT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000132)))

// func RegisterKeyInterrupt()
