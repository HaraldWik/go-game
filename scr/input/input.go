package input

import (
	"runtime"
	"syscall"
)

// *Get array of all pressed keys
func GetPressedKeys() []int {
	switch runtime.GOOS {
	case "windows":
		user32 := syscall.NewLazyDLL("user32.dll")
		keyboard := user32.NewProc("GetAsyncKeyState")

		getKeyState := func(key int) uint16 {
			result, _, _ := keyboard.Call(uintptr(key))
			return uint16(result)
		}

		var pressedKeys []int
		for keycode := 0; keycode < 256; keycode++ {
			if result := getKeyState(keycode); result != 0 {
				pressedKeys = append(pressedKeys, keycode)
			}
		}
		return pressedKeys
	case "darwin", "linux":
		// !No support for your OS
		return nil
	default:
		return nil
	}
}

// *Is interaction
func IsPressed(keycode int) bool {
	keys := GetPressedKeys()
	for _, key := range keys {
		if key == keycode {
			return true
		}
	}
	return false
}

func IsReleased(keycode int) bool {
	return !IsPressed(keycode)
}

// *Key Just interaction
var previousPressedKeyState = map[int]bool{}

func IsJustPressed(keycode int) bool {
	currentState := IsPressed(keycode)
	pressed := currentState && !previousPressedKeyState[keycode]
	previousPressedKeyState[keycode] = currentState

	return pressed
}

var previousReleasedKeyState = map[int]bool{}

func IsJustReleased(keycode int) bool {
	currentState := IsPressed(keycode) && !IsJustPressed(keycode)
	released := !currentState && previousReleasedKeyState[keycode]
	previousReleasedKeyState[keycode] = currentState

	return released
}
