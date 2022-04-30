package color

// PackRGB receives a 16bit color and converts it to 15bit expected by GBA
func PackRGB(r, g, b uint16) uint16 {
	return (r >> 3) | ((g >> 3) << 5) | ((b >> 3) << 10)
}
