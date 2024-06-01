package input

import "runtime"

var (
	EMPTY              int
	KEY_A              int
	KEY_B              int
	KEY_C              int
	KEY_D              int
	KEY_E              int
	KEY_F              int
	KEY_G              int
	KEY_H              int
	KEY_I              int
	KEY_J              int
	KEY_K              int
	KEY_L              int
	KEY_M              int
	KEY_N              int
	KEY_O              int
	KEY_P              int
	KEY_Q              int
	KEY_R              int
	KEY_S              int
	KEY_T              int
	KEY_U              int
	KEY_V              int
	KEY_W              int
	KEY_X              int
	KEY_Y              int
	KEY_Z              int
	KEY_0              int
	KEY_1              int
	KEY_2              int
	KEY_3              int
	KEY_4              int
	KEY_5              int
	KEY_6              int
	KEY_7              int
	KEY_8              int
	KEY_9              int
	KEY_F1             int
	KEY_F2             int
	KEY_F3             int
	KEY_F4             int
	KEY_F5             int
	KEY_F6             int
	KEY_F7             int
	KEY_F8             int
	KEY_F9             int
	KEY_F10            int
	KEY_F11            int
	KEY_F12            int
	KEY_TAB            int
	KEY_ENTER          int
	KEY_SHIFT          int
	KEY_CONTROL        int
	KEY_LEFT_CONTROL   int
	KEY_RIGHT_CONTROL  int
	KEY_ALT            int
	KEY_SPACE          int
	KEY_ESC            int
	KEY_BACKSPACE      int
	KEY_INSERT         int
	KEY_DELETE         int
	KEY_HOME           int
	KEY_END            int
	KEY_PAGEUP         int
	KEY_PAGEDOWN       int
	KEY_LEFT           int
	KEY_UP             int
	KEY_RIGHT          int
	KEY_DOWN           int
	KEY_CAPSLOCK       int
	KEY_NUMLOCK        int
	KEY_SCROLLLOCK     int
	KEY_LEFT_WIN       int
	KEY_RIGHT_WIN      int
	KEY_APPS           int
	KEY_PAUSE          int
	MOUSE_LEFT_BUTTON  int
	MOUSE_RIGHT_BUTTON int
)

func init() {
	switch runtime.GOOS {
	case "windows":
		EMPTY = -1
		KEY_A = 0x41
		KEY_B = 0x42
		KEY_C = 0x43
		KEY_D = 0x44
		KEY_E = 0x45
		KEY_F = 0x46
		KEY_G = 0x47
		KEY_H = 0x48
		KEY_I = 0x49
		KEY_J = 0x4A
		KEY_K = 0x4B
		KEY_L = 0x4C
		KEY_M = 0x4D
		KEY_N = 0x4E
		KEY_O = 0x4F
		KEY_P = 0x50
		KEY_Q = 0x51
		KEY_R = 0x52
		KEY_S = 0x53
		KEY_T = 0x54
		KEY_U = 0x55
		KEY_V = 0x56
		KEY_W = 0x57
		KEY_X = 0x58
		KEY_Y = 0x59
		KEY_Z = 0x5A
		KEY_0 = 0x30
		KEY_1 = 0x31
		KEY_2 = 0x32
		KEY_3 = 0x33
		KEY_4 = 0x34
		KEY_5 = 0x35
		KEY_6 = 0x36
		KEY_7 = 0x37
		KEY_8 = 0x38
		KEY_9 = 0x39
		KEY_F1 = 0x70
		KEY_F2 = 0x71
		KEY_F3 = 0x72
		KEY_F4 = 0x73
		KEY_F5 = 0x74
		KEY_F6 = 0x75
		KEY_F7 = 0x76
		KEY_F8 = 0x77
		KEY_F9 = 0x78
		KEY_F10 = 0x79
		KEY_F11 = 0x7A
		KEY_F12 = 0x7B
		KEY_TAB = 0x09
		KEY_ENTER = 0x0D
		KEY_SHIFT = 0x10
		KEY_CONTROL = 0x11
		KEY_LEFT_CONTROL = 0xA2
		KEY_RIGHT_CONTROL = 0xA3
		KEY_ALT = 0x12
		KEY_SPACE = 0x20
		KEY_ESC = 0x1B
		KEY_BACKSPACE = 0x08
		KEY_INSERT = 0x2D
		KEY_DELETE = 0x2E
		KEY_HOME = 0x24
		KEY_END = 0x23
		KEY_PAGEUP = 0x21
		KEY_PAGEDOWN = 0x22
		KEY_LEFT = 0x25
		KEY_UP = 0x26
		KEY_RIGHT = 0x27
		KEY_DOWN = 0x28
		KEY_CAPSLOCK = 0x14
		KEY_NUMLOCK = 0x90
		KEY_SCROLLLOCK = 0x91
		KEY_LEFT_WIN = 0x5B
		KEY_RIGHT_WIN = 0x5C
		KEY_APPS = 0x5D
		KEY_PAUSE = 0x13
		MOUSE_LEFT_BUTTON = 0x01
		MOUSE_RIGHT_BUTTON = 0x02
	case "darwin", "linux":
		EMPTY = -1
		KEY_A = 0x61
		KEY_B = 0x62
		KEY_C = 0x63
		KEY_D = 0x64
		KEY_E = 0x65
		KEY_F = 0x66
		KEY_G = 0x67
		KEY_H = 0x68
		KEY_I = 0x69
		KEY_J = 0x6A
		KEY_K = 0x6B
		KEY_L = 0x6C
		KEY_M = 0x6D
		KEY_N = 0x6E
		KEY_O = 0x6F
		KEY_P = 0x70
		KEY_Q = 0x71
		KEY_R = 0x72
		KEY_S = 0x73
		KEY_T = 0x74
		KEY_U = 0x75
		KEY_V = 0x76
		KEY_W = 0x77
		KEY_X = 0x78
		KEY_Y = 0x79
		KEY_Z = 0x7A
		KEY_0 = 0x30
		KEY_1 = 0x31
		KEY_2 = 0x32
		KEY_3 = 0x33
		KEY_4 = 0x34
		KEY_5 = 0x35
		KEY_6 = 0x36
		KEY_7 = 0x37
		KEY_8 = 0x38
		KEY_9 = 0x39
		KEY_F1 = 0xFFBE
		KEY_F2 = 0xFFBF
		KEY_F3 = 0xFFC0
		KEY_F4 = 0xFFC1
		KEY_F5 = 0xFFC2
		KEY_F6 = 0xFFC3
		KEY_F7 = 0xFFC4
		KEY_F8 = 0xFFC5
		KEY_F9 = 0xFFC6
		KEY_F10 = 0xFFC7
		KEY_F11 = 0xFFC8
		KEY_F12 = 0xFFC9
		KEY_TAB = 0xFF09
		KEY_ENTER = 0xFF0D
		KEY_SHIFT = 0xFFE1
		KEY_CONTROL = 0xFFE3
		KEY_LEFT_CONTROL = 0xFFE3
		KEY_RIGHT_CONTROL = 0xFFE4
		KEY_ALT = 0xFFE9
		KEY_SPACE = 0x20
		KEY_ESC = 0xFF1B
		KEY_BACKSPACE = 0xFF08
		KEY_INSERT = 0xFF63
		KEY_DELETE = 0xFFFF
		KEY_HOME = 0xFF50
		KEY_END = 0xFF57
		KEY_PAGEUP = 0xFF55
		KEY_PAGEDOWN = 0xFF56
		KEY_LEFT = 0xFF51
		KEY_UP = 0xFF52
		KEY_RIGHT = 0xFF53
		KEY_DOWN = 0xFF54
		KEY_CAPSLOCK = 0xFFE5
		KEY_NUMLOCK = 0xFF7F
		KEY_SCROLLLOCK = 0xFF14
		KEY_LEFT_WIN = 0xFFEB
		KEY_RIGHT_WIN = 0xFFEC
		KEY_APPS = 0xFF67
		KEY_PAUSE = 0xFF13
		MOUSE_LEFT_BUTTON = 0x01
		MOUSE_RIGHT_BUTTON = 0x02
	}

}
