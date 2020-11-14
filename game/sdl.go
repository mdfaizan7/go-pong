package game

import "go-mod/util"

// SetPixel sets the pixel in the texture
func SetPixel(x, y int, c Color, pixels []byte) {
	idx := (y*util.WinWidth + x) * 4

	if idx < len(pixels)-4 && idx >= 0 {
		pixels[idx] = c.R
		pixels[idx+1] = c.G
		pixels[idx+2] = c.B
	}
}

// ConstructPaddle returns a paddle
func ConstructPaddle() *Paddle {
	return &Paddle{Pos{X: 100, Y: 100}, 20, 100, Color{R: 255, G: 255, B: 255}}
}

// Construct Ball returns a ball
func ConstructBall() *Ball {
	return &Ball{Pos{X: 300, Y: 300}, 20, 0, 0, Color{R: 255, G: 255, B: 255}}
}
