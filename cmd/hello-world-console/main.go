package main

import (
	"github.com/gkawamoto/go-libgba/color"
	"github.com/gkawamoto/go-libgba/console"
	"github.com/gkawamoto/go-libgba/display"
)

var (
	red   uint16 = 127
	green uint16 = 127
	blue  uint16 = 127

	frame = 0
)

func main() {
	display.EnableVBlankInterrupts()
	display.RegisterVBlankUpdate(update)
	display.SetDisplayOptions(display.Mode3, display.BG2)

	select {}
}

func update() {
	frame++
	if frame%10 == 0 {
		console.Println("hey lelinha :)")
		console.Println("hello world!")

		// Coloring the world!
		red += 50
		if red >= 255 {
			red = 127
			green += 50
		}
		if green >= 255 {
			green = 127
			blue += 50
		}
		if blue >= 255 {
			blue = 127
		}
		console.SetMode3ForegroundColor(color.PackRGB(red+127, green+127, blue+127))
	}
}
