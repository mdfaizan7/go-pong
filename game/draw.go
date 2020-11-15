package game

import "go-mod/util"

// Draw function draws the updated pixels of Paddle to the screen as texture
func (p *Paddle) Draw(pixels []byte) {
	startX, startY := p.X-p.W/2, p.Y-p.H/2

	for y := 0; y < int(p.H); y++ {
		for x := 0; x < int(p.W); x++ {
			SetPixel(startX+float32(x), startY+float32(y), p.Color, pixels)
		}
	}

	numX := util.Lerp(p.X, GetCenter().X, 0.2)
	DrawNumber(Pos{numX, 35}, p.Color, 10, p.Score, pixels)
}

// Draw function draws the updated pixels of Ball to the screen as texture
func (b *Ball) Draw(pixels []byte) {
	for y := -b.Radius; y < b.Radius; y++ {
		for x := -b.Radius; x < b.Radius; x++ {
			if x*x+y*y < b.Radius*b.Radius {
				SetPixel(b.X+x, b.Y+y, b.Color, pixels)
			}
		}
	}
}
