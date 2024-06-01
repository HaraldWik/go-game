package win

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Window struct {
	width     int
	height    int
	minWidth  int
	minHeight int
	maxWidth  int
	maxHeight int

	name string

	flags uint32

	SDL *sdl.Window
}

func NewWindow(width, height int, name string) Window {
	return Window{
		width:  width,
		height: height,
		name:   name,
	}
}

func NewWindowAdvanced(width, height, minWidth, minHeight, maxWidth, maxHeight int, name string, flags uint32) Window {
	return Window{
		width:     width,
		height:    height,
		minWidth:  minWidth,
		minHeight: minHeight,
		maxWidth:  maxWidth,
		maxHeight: maxHeight,
		name:      name,
		flags:     flags,
	}
}
