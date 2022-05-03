package main

import (
	"log"

	"github.com/gkawamoto/go-libgba/agbprint"
	"github.com/gkawamoto/go-libgba/display"
)

func main() {
	agbprint.Init()

	display.RegisterVBlankUpdate(update)
	display.EnableVBlankInterrupts()

	select {}
}

var index int

func update() {
	log.Printf("hello world! %d", index)
	index++
}
