package window

import (
	"log"

	err_ "github.com/HaraldWik/go-game/scr/err"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/veandco/go-sdl2/sdl"
)

func (window *Window) InitOpenGLRenderer() sdl.GLContext {
	//* Initialize SDL video subsystem
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		log.Fatalln(err_.FAILED_INIT+"SDL video:", err)
	}

	//* Set OpenGL attributes
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 2)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 1)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)

	//* Create OpenGL context
	glContext, err := window.SDL.GLCreateContext()
	if err != nil {
		log.Fatalln(err_.FAILED_INIT+"OpenGL context for window:", err)
	}

	//* Initialize OpenGL
	if err := gl.Init(); err != nil {
		log.Fatalln(err_.FAILED_INIT, "OpenGL", err)
	}

	//* Enable depth test
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LEQUAL)

	return glContext
}

func (window Window) BeginDrawOpenGL() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (window Window) EndDrawOpenGL() {
	window.SDL.GLSwap()
}
