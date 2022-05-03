package console

import (
	"fmt"
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

var (
	mode3ForegroundColor = color.PackRGB(255, 255, 255)
	mode3BackgroundColor = color.PackRGB(0, 0, 0)

	mode4PaletteBackgroundIndex uint8 = 0
	mode4PaletteForegroundIndex uint8 = 1
)

type Mode uint8

var (
	Mode3 Mode = 0
	Mode4 Mode = 1
)

var mode Mode = Mode3
var frontPage = true

// Println writes the text on screen and move to the next line
func Println(text string) {
	Print(text)
	currentColumn = 0
	currentLine++
}

// Print writes the text on screen.
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
		if currentLine >= maxLines {
			currentLine = 0
		}
	}
}

// Printf is a syntactic sugar for Print(fmt.Sprintf(format, a...))
func Printf(format string, a ...interface{}) {
	Print(fmt.Sprintf(format, a...))
}

// SetMode3ForegroundColor sets the 15bit color used to print the foreground of the characters. Use color.PackRGB to create a nice 15bit color.
func SetMode3ForegroundColor(value uint16) {
	mode3ForegroundColor = value
}

// SetMode3BackgroundColor sets the 15bit color used to print the background of the characters. Use color.PackRGB to create a nice 15bit color.
func SetMode3BackgroundColor(value uint16) {
	mode3BackgroundColor = value
}

// SetMode4PaletteForegroundIndex sets which index of the defined palette will be used for foreground color.
func SetMode4PaletteForegroundIndex(value uint8) {
	mode4PaletteForegroundIndex = value
}

// SetMode4PaletteBackgroundIndex sets which index of the defined palette will be used for background color.
func SetMode4PaletteBackgroundIndex(value uint8) {
	mode4PaletteBackgroundIndex = value
}

// SetMode changes the mode used to render console characters. Match this with the mode used in `display.SetDisplayOptions`
func SetMode(value Mode) {
	mode = value
}

// Reset sets the current line & column back to 0. This function DOES NOT clear up the screen, this is up to you to do so.
func Reset() {
	currentLine = 0
	currentColumn = 0
}

// SetPage defines which page will be used to draw the characters
func SetPage(value bool) {
	frontPage = value
}

// draw draws the character pixel by pixel in the specified line/column
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
					mode3.SetPixel(x+column*(fontWidth+1), y+line*fontHeight, mode3ForegroundColor)
				} else {
					mode3.SetPixel(x+column*(fontWidth+1), y+line*fontHeight, mode3BackgroundColor)
				}
			case Mode4:
				if value {
					mode4.SetPixel(x+column*(fontWidth+1), y+line*fontHeight, mode4PaletteForegroundIndex, frontPage)
				} else {
					mode4.SetPixel(x+column*(fontWidth+1), y+line*fontHeight, mode4PaletteBackgroundIndex, frontPage)
				}
			}
		}
	}
}
