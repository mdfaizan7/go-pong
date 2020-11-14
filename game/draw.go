package game

// Draw function draws the updated pixels of Paddle to the screen as texture
func (p *Paddle) Draw(pixels []byte) {
	startX, startY := int(p.X)-p.W/2, int(p.Y)-p.H/2

	for y := 0; y < p.H; y++ {
		for x := 0; x < p.W; x++ {
			SetPixel(startX+x, startY+y, p.Color, pixels)
		}
	}
}

// Draw function draws the updated pixels of Ball to the screen as texture
func (b *Ball) Draw(pixels []byte) {
	for y := -b.Radius; y < b.Radius; y++ {
		for x := -b.Radius; x < b.Radius; x++ {
			if x*x+y*y < b.Radius*b.Radius {
				SetPixel(int(b.X)+x, int(b.Y)+y, b.Color, pixels)
			}
		}
	}
}
