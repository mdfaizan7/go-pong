package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	window, err := sdl.CreateWindow("Testing SDL2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}
	defer window.Destroy()

	sdl.Delay(2000)
}
