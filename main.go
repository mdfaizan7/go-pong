package main

import (
	"go-mod/game"
	"go-mod/util"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		log.Fatal(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Testing SDL2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
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
			game.SetPixel(x, y,
				util.Color{
					R: 255,
					G: 0,
					B: 0,
				}, pixels)
		}
	}
	tex.Update(nil, pixels, util.WinWidth*4)
	renderer.Copy(tex, nil, nil)
	renderer.Present()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
	}
}
