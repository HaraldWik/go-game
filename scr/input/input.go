package input

import (
	"runtime"
	"syscall"

	err_ "github.com/HaraldWik/go-game/scr/err"
)

// *Get array of all pressed keys
func GetPressedKeys() []int {
	switch runtime.GOOS {
	case "windows":
		user32 := syscall.NewLazyDLL("user32.dll")
		procGetAsyncKeyState := user32.NewProc("GetAsyncKeyState")

		getKey := func(key int) uint16 {
			result, _, _ := procGetAsyncKeyState.Call(uintptr(key))
			return uint16(result)
		}

		var pressedKeys []int
		for keycode := 0; keycode < 256; keycode++ {
			if result := getKey(keycode); result != 0 {
				pressedKeys = append(pressedKeys, keycode)
			}
		}

		return pressedKeys
	case "darwin", "linux":
		println(err_.WARNING + "No input support for MacOS & Linux! *COMING SOON*")
		// !No support for MacOS & Linux yet! *COMING SOON*
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
var previousPressed = map[int]bool{}

func IsJustPressed(keycode int) bool {
	current := IsPressed(keycode)
	pressed := current && !previousPressed[keycode]
	previousPressed[keycode] = current

	return pressed
}

var previousReleased = map[int]bool{}

func IsJustReleased(keycode int) bool {
	current := IsPressed(keycode) && !IsJustPressed(keycode)
	released := !current && previousReleased[keycode]
	previousReleased[keycode] = current

	return released
}
