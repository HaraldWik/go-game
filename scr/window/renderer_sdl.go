package win

import (
	"log"

	err_ "github.com/HaraldWik/go-game/scr/err"

	"github.com/veandco/go-sdl2/sdl"
)

func (window *Window) InitSDLRenderer() *sdl.Renderer {
	renderer, err := sdl.CreateRenderer(window.SDL, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatalln(err_.FAILED_TO_INIT+"SDL renderer", err)
		renderer.Destroy()
	}

	return renderer
}
