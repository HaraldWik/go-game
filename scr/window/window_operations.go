package win

import (
	"log"
	"runtime"
	"syscall"

	err_ "github.com/HaraldWik/go-game/scr/err"

	"github.com/veandco/go-sdl2/sdl"
)

// * Opens window
func (window *Window) Open() {
	runtime.LockOSThread()

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatalf(err_.FAILED_TO_INIT+"SDL: %v", err)
	}
	defer runtime.UnlockOSThread()

	var err error
	window.SDL, err = sdl.CreateWindow(
		window.name,
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(window.width), int32(window.height),
		window.flags,
	)
	if err != nil {
		log.Fatalf(err_.FAILED_TO_INIT+"window: %v", err)
	}
}

// * Closes window
func (window Window) Close() {
	window.SDL.Destroy()
}

// * Closes window on window exit
func (window Window) CloseOnExit() {
	switch window.GetSDLEvents().(type) {
	case *sdl.QuitEvent:
		syscall.Exit(0)
	}
}
