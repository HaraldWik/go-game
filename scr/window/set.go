package window

//* Set size
func (window *Window) SetSize(width, height int) {
	window.width = width
	window.height = height

	window.SDL.SetSize(int32(width), int32(height))
}

func (window *Window) SetSizeWidth(width int) {
	window.width = width

	window.SDL.SetSize(int32(width), int32(window.height))
}

func (window *Window) SetSizeHeight(height int) {
	window.height = height

	window.SDL.SetSize(int32(window.width), int32(height))
}

//* Set min size
func (window *Window) SetMinSize(width, height int) {
	window.minWidth = width
	window.minHeight = height

	window.SDL.SetMinimumSize(int32(width), int32(height))
}

func (window *Window) SetMinSizeWidth(width int) {
	window.minWidth = width

	window.SDL.SetSize(int32(width), int32(window.minHeight))
}

func (window *Window) SetMinSizeHeight(height int) {
	window.minHeight = height

	window.SDL.SetSize(int32(window.minWidth), int32(height))
}

//* Set max size
func (window *Window) SetMaxSize(width, height int) {
	window.maxWidth = width
	window.maxHeight = height

	window.SDL.SetMaximumSize(int32(width), int32(height))

}

func (window *Window) SetMaxSizeWidth(width int) {
	window.maxWidth = width

	window.SDL.SetMaximumSize(int32(width), int32(window.maxHeight))
}

func (window *Window) SetMaxSizeHeight(height int) {
	window.maxHeight = height

	window.SDL.SetMaximumSize(int32(window.maxWidth), int32(height))
}

//* Set name
func (window *Window) SetName(name string) {
	window.name = name

	window.SDL.SetTitle(name)
}

//* Set flags
func (window *Window) SetFlags(flags uint32) {
	window.flags = flags
}
