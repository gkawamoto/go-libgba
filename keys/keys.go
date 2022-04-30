package keys

import (
	"runtime/volatile"
	"unsafe"
)

var (
	// Register pointing to the memory reference of the buttons controller (0400:0130h)
	_REG_KEYINPUT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000130)))

	_KEY_A      uint16 = 0x0001
	_KEY_B      uint16 = 0x0002
	_KEY_SELECT uint16 = 0x0004
	_KEY_START  uint16 = 0x0008
	_KEY_RIGHT  uint16 = 0x0010
	_KEY_LEFT   uint16 = 0x0020
	_KEY_UP     uint16 = 0x0040
	_KEY_DOWN   uint16 = 0x0080
	_KEY_R      uint16 = 0x0100
	_KEY_L      uint16 = 0x0200
)

type State struct {
	keys uint16
}

// CurrentState returns an object containing the current state of buttons in GBA.
func CurrentState() *State {
	keys := _REG_KEYINPUT.Get()
	return &State{keys}
}

// IsKeyADown returns true if the key A is currently being held.
func (s *State) IsKeyADown() bool {
	return ^s.keys&_KEY_A == _KEY_A
}

// IsKeyBDown returns true if the key B is currently being held.
func (s *State) IsKeyBDown() bool {
	return ^s.keys&_KEY_B == _KEY_B
}

// IsKeyDownDown returns true if the key DOWN is currently being held.
func (s *State) IsKeyDownDown() bool {
	return ^s.keys&_KEY_DOWN == _KEY_DOWN
}

// IsKeyUpDown returns true if the key UP is currently being held.
func (s *State) IsKeyUpDown() bool {
	return ^s.keys&_KEY_UP == _KEY_UP
}

// IsKeyLeftDown returns true if the key LEFT is currently being held.
func (s *State) IsKeyLeftDown() bool {
	return ^s.keys&_KEY_LEFT == _KEY_LEFT
}

// IsKeyRightDown returns true if the key RIGHT is currently being held.
func (s *State) IsKeyRightDown() bool {
	return ^s.keys&_KEY_RIGHT == _KEY_RIGHT
}

// IsKeyLDown returns true if the key L is currently being held.
func (s *State) IsKeyLDown() bool {
	return ^s.keys&_KEY_L == _KEY_L
}

// IsKeyRDown returns true if the key R is currently being held.
func (s *State) IsKeyRDown() bool {
	return ^s.keys&_KEY_R == _KEY_R
}

// IsKeyStartDown returns true if the key START is currently being held.
func (s *State) IsKeyStartDown() bool {
	return ^s.keys&_KEY_START == _KEY_START
}

// IsKeySelectDown returns true if the key SELECT is currently being held.
func (s *State) IsKeySelectDown() bool {
	return ^s.keys&_KEY_SELECT == _KEY_SELECT
}
