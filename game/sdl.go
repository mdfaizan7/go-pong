package game

import "go-mod/util"

// SetPixel sets the pixel in the texture
func SetPixel(x, y int, c util.Color, pixels []byte) {
	idx := (y*util.WinWidth + x) * 4

	if idx < len(pixels)-4 && idx >= 0 {
		pixels[idx] = c.R
		pixels[idx+1] = c.G
		pixels[idx+2] = c.B
	}
}
