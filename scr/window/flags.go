package window

const (
	FLAG_RESIZABLE          = 0x00000020 //* Window can be resized
	FLAG_FULLSCREEN         = 0x00000001 //* Window fullscreen
	FLAG_FULLSCREEN_DESKTOP = 0x00001000 //* Window fullscreen but worse
	FLAG_MINIMIZED          = 0x00000040 //* Window is minimized
	FLAG_MAXIMIZED          = 0x00000080 //* Window is maximized
	FLAG_BORDERLESS         = 0x00000010 //* Window no decoration
	FLAG_SHOW               = 0x00000004 //* Window is visible
	FLAG_HIDE               = 0x00000008 //* Window is not visible
	FLAG_FOREIGN            = 0x00000800 //* Window not created by SDL
	FLAG_ALLOW_HIGHDPI      = 0x00002000 //* Window should be created in high-DPI mode if supported.
	FLAG_ALWAYS_ON_TOP      = 0x00008000 //* Window should always be above other windows
	FLAG_SKIP_TASKBAR       = 0x00010000 //* Window should not be added to the taskbar
	FLAG_UTILITY            = 0x00020000 //* Window should be treated as a utility
	FLAG_TOOLTIP            = 0x00040000 //* Window should be treated as a tooltip
	FLAG_POPUP_MENU         = 0x00080000 //* Window should be treated as a popup menu
	FLAG_OPENGL             = 0x00000002 //* Window usable with OpenGL context
	FLAG_VULKAN             = 0x10000000 //* Window usable for Vulkan surface
	FLAG_METAL              = 0x20000000 //* Window usable for Metal view
)
