package main

import (
	"go-mod/game"
	"go-mod/util"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

// --
type gameState int

const (
	start gameState = iota
	play
)

var state = start

// SetStateStart sets the game state to start
func SetStateStart() {
	state = start
}

// --

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		log.Fatal(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Go Pong", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(util.WinWidth), int32(util.WinHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatal(err)
	}
	defer renderer.Destroy()

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(util.WinWidth), int32(util.WinHeight))
	if err != nil {
		log.Fatal(err)
	}
	defer tex.Destroy()

	pixels := make([]byte, util.WinWidth*util.WinHeight*4)

	for y := 0; y < util.WinHeight; y++ {
		for x := 0; x < util.WinWidth; x++ {
			game.SetPixel(float32(x), float32(y),
				game.Color{
					R: 0,
					G: 0,
					B: 0,
				}, pixels)
		}
	}

	player1 := game.ConstructPaddle(50, 100)
	player2 := game.ConstructPaddle(float32(util.WinWidth-50), 100)
	ball := game.ConstructBall()

	keyState := sdl.GetKeyboardState()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		if state == play {
			// Update functions
			player1.Update(keyState)
			player2.AIUpdate(ball)
			ball.Update(player1, player2, SetStateStart)
		} else if state == start {
			if keyState[sdl.SCANCODE_SPACE] != 0 {
				if player1.Score == 5 || player2.Score == 5 {
					player1.Score = 0
					player2.Score = 0
				}
				state = play
			}
		}

		util.ClearPixels(pixels)
		// Draw functions
		player1.Draw(pixels)
		ball.Draw(pixels)
		player2.Draw(pixels)

		tex.Update(nil, pixels, util.WinWidth*4)
		renderer.Copy(tex, nil, nil)
		renderer.Present()

		sdl.Delay(15)
	}
}
