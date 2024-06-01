package main

import (
	"go-game/scr/input"
)

func main() {
	for {
		if input.IsJustPressed(input.KEY_ESC) {
			break
		}

		if input.IsJustReleased(input.KEY_H) {
			print()
		}
	}
}
