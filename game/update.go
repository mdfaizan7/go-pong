package game

import (
	"go-mod/util"

	"github.com/veandco/go-sdl2/sdl"
)

type function func()

// Update function updates the position of Ball
func (b *Ball) Update(p1 *Paddle, p2 *Paddle, setStateStart function) {
	b.X += b.XV
	b.Y += b.YV

	if b.Y-b.Radius < 0 || b.Y+b.Radius > float32(util.WinHeight) {
		b.YV = -b.YV
	}

	if b.X < 0 {
		p2.Score++
		b.Pos = GetCenter()
		setStateStart()
	} else if int(b.X) > util.WinWidth {
		p1.Score++
		b.Pos = GetCenter()
		setStateStart()
	}

	if b.X < p1.X+p1.W/2 && // check if paddle X is equal to ball X
		b.Y > p1.Y-p1.H/2 && b.Y < p1.Y+p1.H/2 { // check if the paddle covers same height as the ball
		b.XV = -b.XV
	}

	if b.X > p2.X-p2.W/2 &&
		b.Y > p2.Y-p2.H/2 && b.Y < p2.Y+p2.H/2 {
		b.XV = -b.XV
	}
}

// Update function updates the position of paddle
func (p *Paddle) Update(keyState []uint8) {
	if keyState[sdl.SCANCODE_UP] != 0 {
		p.Y -= 10
	}

	if keyState[sdl.SCANCODE_DOWN] != 0 {
		p.Y += 10
	}
}

// AIUpdate function updates the position of AI paddle
func (p *Paddle) AIUpdate(b *Ball) {
	p.Y = b.Y
}

// GetCenter returns the position of center of screen
func GetCenter() Pos {
	return Pos{float32(util.WinWidth / 2), float32(util.WinHeight / 2)}
}
