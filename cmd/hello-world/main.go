package main

import (
	"github.com/gkawamoto/go-libgba/console"
	"github.com/gkawamoto/go-libgba/display"
)

func main() {
	display.EnableVBlankInterrupts()
	display.SetDisplayOptions(display.Mode3, display.BG2)
	console.Println("hey lelinha :)")
	console.Println("hello world!")

	select {}
}
