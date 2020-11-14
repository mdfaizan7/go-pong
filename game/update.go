package game

import (
	"go-mod/util"

	"github.com/veandco/go-sdl2/sdl"
)

// Update function updates the position of Ball
func (b *Ball) Update() {
	b.X += b.XV
	b.Y += b.YV

	if int(b.Y)-b.Radius < 0 || int(b.Y)+b.Radius > util.WinHeight {
		b.YV = -b.YV
	}

	if b.X < 0 || int(b.X) > util.WinWidth {
		b.X, b.Y = 300, 300 // To be changed
	}
}

// Update function updates the position of paddle
func (p *Paddle) Update(keyState []uint8) {
	if keyState[sdl.SCANCODE_UP] != 0 {
		p.Y--
	}

	if keyState[sdl.SCANCODE_DOWN] != 0 {
		p.Y++
	}
}
