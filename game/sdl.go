package game

import "go-mod/util"

// SetPixel sets the pixel in the texture
func SetPixel(x, y float32, c Color, pixels []byte) {
	idx := int(y*float32(util.WinWidth)+x) * 4

	if idx < len(pixels)-4 && idx >= 0 {
		pixels[idx] = c.R
		pixels[idx+1] = c.G
		pixels[idx+2] = c.B
	}
}

// ConstructPaddle returns a paddle
func ConstructPaddle(X, Y float32) *Paddle {
	return &Paddle{Pos{X: X, Y: Y}, 20, 100, Color{R: 255, G: 255, B: 255}, 0}
}

// ConstructBall returns a ball
func ConstructBall() *Ball {
	return &Ball{Pos{X: 300, Y: 300}, 20, 10, 10, Color{R: 255, G: 255, B: 255}}
}

// DrawNumber draws the numbers on screen
func DrawNumber(pos Pos, color Color, size int, num int, pixels []byte) {
	startX := int(pos.X) - (size*3)/2
	startY := int(pos.Y) - (size*5)/2

	for i, v := range util.Nums[num] {
		if v == 1 {
			for y := startY; y < startY+size; y++ {
				for x := startX; x < startX+size; x++ {
					SetPixel(float32(x), float32(y), color, pixels)
				}
			}
		}

		startX += size

		if (i+1)%3 == 0 {
			startY += size
			startX -= size * 3
		}
	}
}
