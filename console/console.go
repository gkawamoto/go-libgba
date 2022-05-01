package console

import (
	"strings"

	"github.com/gkawamoto/go-libgba/color"
	"github.com/gkawamoto/go-libgba/display/mode3"
	"github.com/gkawamoto/go-libgba/display/mode4"
)

var (
	currentLine   uint16 = 0
	currentColumn uint16 = 0
)

const (
	maxColumns uint16 = 240 / 5
	maxLines   uint16 = 160 / 7
)

var mode3ForegroundColor = color.PackRGB(255, 255, 255)
var mode3BackgroundColor = color.PackRGB(0, 0, 0)
var mode4PaletteIndex uint8 = 1

type Mode uint8

var (
	Mode3 Mode = 0
	Mode4 Mode = 1
)

var mode Mode = Mode3
var frontPage = true

func Println(text string) {
	Print(text)
	currentColumn = 0
	currentLine++
}

func Print(text string) {
	text = strings.ToLower(text)
	for _, c := range text {
		if c == '\n' {
			currentColumn = 0
			currentLine++
			continue
		} else if c != ' ' {
			draw(c, currentLine, currentColumn)
		}
		currentColumn++
		if currentColumn >= maxColumns {
			currentColumn = 0
			currentLine++
		}
	}
}

func SetMode3ForegroundColor(value uint16) {
	mode3ForegroundColor = value
}

func SetMode3BackgroundColor(value uint16) {
	mode3BackgroundColor = value
}

func SetMode4PaletteIndex(value uint8) {
	mode4PaletteIndex = value
}

func SetMode(value Mode) {
	mode = value
}

func SetPage(value bool) {
	frontPage = value
}

func draw(char rune, line, column uint16) {
	pos, ok := runes[char]
	if !ok {
		pos = runes['?']
	}

	var x, y uint16
	for y = 0; y < fontHeight; y++ {
		for x = 0; x < fontWidth; x++ {
			value := font[(x+fontWidth*charCount*y)+(pos*fontWidth)]
			switch mode {
			case Mode3:
				if value {
					mode3.SetPixel(uint8(x+column*(fontWidth+1)), uint8(y+line*fontHeight), mode3ForegroundColor)
				} else {
					mode3.SetPixel(uint8(x+column*(fontWidth+1)), uint8(y+line*fontHeight), mode3BackgroundColor)
				}
			case Mode4:
				if value {
					mode4.SetPixel(uint8(x+column*(fontWidth+1)), uint8(y+line*fontHeight), mode4PaletteIndex, frontPage)
				}
			}
		}
	}
}
