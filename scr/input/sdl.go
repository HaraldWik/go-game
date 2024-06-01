package input

import (
	"github.com/veandco/go-sdl2/sdl"
)

// GetPressedKeys returns an array of currently pressed keys
func GetPressedKeys() []sdl.Keycode {
	sdl.Init(sdl.INIT_EVENTS) // Initialize only the event subsystem

	var pressedKeys []sdl.Keycode

	for _, key := range sdl.GetKeyboardState() {
		if key == 1 {
			pressedKeys = append(pressedKeys, sdl.Keycode(len(pressedKeys)))
		}
	}

	sdl.QuitSubSystem(sdl.INIT_EVENTS) // Quit event subsystem before returning

	return pressedKeys
}

func main() {
	pressedKeys := GetPressedKeys()

	for _, key := range pressedKeys {
		println("Key pressed:", key)
	}
}
