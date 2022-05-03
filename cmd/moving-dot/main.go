package main

import (
	"github.com/gkawamoto/go-libgba/color"
	"github.com/gkawamoto/go-libgba/console"
	"github.com/gkawamoto/go-libgba/display"
	"github.com/gkawamoto/go-libgba/display/mode3"
	"github.com/gkawamoto/go-libgba/display/mode4"
	"github.com/gkawamoto/go-libgba/keys"
)

var (
	x uint16 = 30
	y uint16 = 7
)
var redPaletteIndex uint8 = 1

var mode = 4

func main() {
	display.EnableVBlankInterrupts()
	display.RegisterVBlankUpdate(update)

	if mode == 4 {
		initMode4()
		mode4.SetPixel(x, y, redPaletteIndex, true)
	} else if mode == 3 {
		initMode3()
		mode3.SetPixel(x, y, color.PackRGB(255, 0, 0))
	}

	select {}
}

func initMode4() {
	display.SetDisplayOptions(display.Mode4, display.BG2)
	mode4.SetPalette(redPaletteIndex, color.PackRGB(255, 0, 0))
	console.SetMode(console.Mode4)
	console.Reset()
	console.Println("Mode 4. Press L to go to Mode 3")
}

func initMode3() {
	display.SetDisplayOptions(display.Mode3, display.BG2)
	console.SetMode(console.Mode3)
	console.Reset()
	console.Println("Mode 3. Press R to go to Mode 4")
}

func update() {
	keys := keys.CurrentState()

	newX := x
	newY := y

	if keys.IsKeyLeftDown() {
		newX--
		if newX <= 1 {
			newX = 1
		}
	}

	if keys.IsKeyRightDown() {
		newX++
		if newX >= 239 {
			newX = 239
		}
	}

	if keys.IsKeyUpDown() {
		newY--
		if newY <= 1 {
			newY = 1
		}
	}

	if keys.IsKeyDownDown() {
		newY++
		if newY >= 159 {
			newY = 159
		}
	}

	if keys.IsKeyLDown() {
		mode = 3
		initMode3()
	}
	if keys.IsKeyRDown() {
		mode = 4
		initMode4()
	}

	// Found changes, lets render them
	if newX != x || newY != y {
		if mode == 4 {
			mode4.SetPixel(x, y, 0, true)                     // Clear old one
			mode4.SetPixel(newX, newY, redPaletteIndex, true) // Render new one
		} else if mode == 3 {
			mode3.SetPixel(x, y, color.PackRGB(0, 0, 0))         // Clear old one
			mode3.SetPixel(newX, newY, color.PackRGB(255, 0, 0)) // Render new one
		}

		// Update state so we can detect new changes
		x = newX
		y = newY
	}
}
