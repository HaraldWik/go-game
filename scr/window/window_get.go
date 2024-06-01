package win

import "github.com/veandco/go-sdl2/sdl"

//* Get size
func (window *Window) GetSize() (int, int) {
	width, height := window.SDL.GetSize()

	return int(width), int(height)
}

func (window *Window) GetSizeWidth() int {
	width, _ := window.SDL.GetSize()

	return int(width)
}

func (window *Window) GetSizeHeight() int {
	_, height := window.SDL.GetSize()

	return int(height)
}

//* Get min size
func (window *Window) GetMinSize() (int, int) {
	width, height := window.SDL.GetMinimumSize()

	return int(width), int(height)
}

func (window *Window) GetMinSizeWidth() int {
	width, _ := window.SDL.GetMinimumSize()

	return int(width)
}

func (window *Window) GetMinSizeHeight() int {
	_, height := window.SDL.GetMinimumSize()

	return int(height)
}

//* Get max size
func (window *Window) GetMaxSize() (int, int) {
	width, height := window.SDL.GetMaximumSize()

	return int(width), int(height)
}

func (window *Window) GetSizeMaxWidth() int {
	width, _ := window.SDL.GetMaximumSize()

	return int(width)
}

func (window *Window) GetSizeMaxHeight() int {
	_, height := window.SDL.GetMaximumSize()

	return int(height)
}

//* Get name
func (window *Window) GetName() string {
	return window.name
}

//* Get flags
func (window *Window) GetFlags() uint32 {
	return window.flags
}

//* Gets SDL events
func (window Window) GetSDLEvents() interface{} {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		return event
	}

	return nil
}

//* Is window active
func (window Window) IsActive() bool {
	switch window.GetSDLEvents().(type) {
	case *sdl.QuitEvent:
		return true
	case *sdl.WindowEvent:
		return true
	case *sdl.KeyboardEvent:
		return true
	case *sdl.MouseButtonEvent:
		return true
	case *sdl.MouseMotionEvent:
		return true
	case *sdl.MouseWheelEvent:
		return true
	case *sdl.JoyButtonEvent:
		return true
	case *sdl.JoyAxisEvent:
		return true
	case *sdl.ControllerButtonEvent:
		return true
	case *sdl.ControllerAxisEvent:
		return true
	case *sdl.ControllerDeviceEvent:
		return true
	case *sdl.TouchFingerEvent:
		return true
	case *sdl.MultiGestureEvent:
		return true
	case *sdl.DollarGestureEvent:
		return true
	case *sdl.DropEvent:
		return true
	case *sdl.ClipboardEvent:
		return true
	default:
		return false
	}
}
