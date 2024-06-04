package input

import "runtime"

var (
	EMPTY              int
	K_A                int
	K_B                int
	K_C                int
	K_D                int
	K_E                int
	K_F                int
	K_G                int
	K_H                int
	K_I                int
	K_J                int
	K_K                int
	K_L                int
	K_M                int
	K_N                int
	K_O                int
	K_P                int
	K_Q                int
	K_R                int
	K_S                int
	K_T                int
	K_U                int
	K_V                int
	K_W                int
	K_X                int
	K_Y                int
	K_Z                int
	K_0                int
	K_1                int
	K_2                int
	K_3                int
	K_4                int
	K_5                int
	K_6                int
	K_7                int
	K_8                int
	K_9                int
	K_F1               int
	K_F2               int
	K_F3               int
	K_F4               int
	K_F5               int
	K_F6               int
	K_F7               int
	K_F8               int
	K_F9               int
	K_F10              int
	K_F11              int
	K_F12              int
	K_TAB              int
	K_ENTER            int
	K_SHIFT            int
	K_CONTROL          int
	K_LEFT_CONTROL     int
	K_RIGHT_CONTROL    int
	K_ALT              int
	K_SPACE            int
	K_ESC              int
	K_BACKSPACE        int
	K_INSERT           int
	K_DELETE           int
	K_HOME             int
	K_END              int
	K_PAGEUP           int
	K_PAGEDOWN         int
	K_LEFT             int
	K_UP               int
	K_RIGHT            int
	K_DOWN             int
	K_CAPSLOCK         int
	K_NUMLOCK          int
	K_SCROLLLOCK       int
	K_LEFT_WIN         int
	K_RIGHT_WIN        int
	K_APPS             int
	K_PAUSE            int
	MOUSE_LEFT_BUTTON  int
	MOUSE_RIGHT_BUTTON int
)

func init() {
	switch runtime.GOOS {
	case "windows":
		EMPTY = -1
		K_A = 0x41
		K_B = 0x42
		K_C = 0x43
		K_D = 0x44
		K_E = 0x45
		K_F = 0x46
		K_G = 0x47
		K_H = 0x48
		K_I = 0x49
		K_J = 0x4A
		K_K = 0x4B
		K_L = 0x4C
		K_M = 0x4D
		K_N = 0x4E
		K_O = 0x4F
		K_P = 0x50
		K_Q = 0x51
		K_R = 0x52
		K_S = 0x53
		K_T = 0x54
		K_U = 0x55
		K_V = 0x56
		K_W = 0x57
		K_X = 0x58
		K_Y = 0x59
		K_Z = 0x5A
		K_0 = 0x30
		K_1 = 0x31
		K_2 = 0x32
		K_3 = 0x33
		K_4 = 0x34
		K_5 = 0x35
		K_6 = 0x36
		K_7 = 0x37
		K_8 = 0x38
		K_9 = 0x39
		K_F1 = 0x70
		K_F2 = 0x71
		K_F3 = 0x72
		K_F4 = 0x73
		K_F5 = 0x74
		K_F6 = 0x75
		K_F7 = 0x76
		K_F8 = 0x77
		K_F9 = 0x78
		K_F10 = 0x79
		K_F11 = 0x7A
		K_F12 = 0x7B
		K_TAB = 0x09
		K_ENTER = 0x0D
		K_SHIFT = 0x10
		K_CONTROL = 0x11
		K_LEFT_CONTROL = 0xA2
		K_RIGHT_CONTROL = 0xA3
		K_ALT = 0x12
		K_SPACE = 0x20
		K_ESC = 0x1B
		K_BACKSPACE = 0x08
		K_INSERT = 0x2D
		K_DELETE = 0x2E
		K_HOME = 0x24
		K_END = 0x23
		K_PAGEUP = 0x21
		K_PAGEDOWN = 0x22
		K_LEFT = 0x25
		K_UP = 0x26
		K_RIGHT = 0x27
		K_DOWN = 0x28
		K_CAPSLOCK = 0x14
		K_NUMLOCK = 0x90
		K_SCROLLLOCK = 0x91
		K_LEFT_WIN = 0x5B
		K_RIGHT_WIN = 0x5C
		K_APPS = 0x5D
		K_PAUSE = 0x13
		MOUSE_LEFT_BUTTON = 0x01
		MOUSE_RIGHT_BUTTON = 0x02
	case "darwin", "linux":
		EMPTY = -1
		K_A = 0x61
		K_B = 0x62
		K_C = 0x63
		K_D = 0x64
		K_E = 0x65
		K_F = 0x66
		K_G = 0x67
		K_H = 0x68
		K_I = 0x69
		K_J = 0x6A
		K_K = 0x6B
		K_L = 0x6C
		K_M = 0x6D
		K_N = 0x6E
		K_O = 0x6F
		K_P = 0x70
		K_Q = 0x71
		K_R = 0x72
		K_S = 0x73
		K_T = 0x74
		K_U = 0x75
		K_V = 0x76
		K_W = 0x77
		K_X = 0x78
		K_Y = 0x79
		K_Z = 0x7A
		K_0 = 0x30
		K_1 = 0x31
		K_2 = 0x32
		K_3 = 0x33
		K_4 = 0x34
		K_5 = 0x35
		K_6 = 0x36
		K_7 = 0x37
		K_8 = 0x38
		K_9 = 0x39
		K_F1 = 0xFFBE
		K_F2 = 0xFFBF
		K_F3 = 0xFFC0
		K_F4 = 0xFFC1
		K_F5 = 0xFFC2
		K_F6 = 0xFFC3
		K_F7 = 0xFFC4
		K_F8 = 0xFFC5
		K_F9 = 0xFFC6
		K_F10 = 0xFFC7
		K_F11 = 0xFFC8
		K_F12 = 0xFFC9
		K_TAB = 0xFF09
		K_ENTER = 0xFF0D
		K_SHIFT = 0xFFE1
		K_CONTROL = 0xFFE3
		K_LEFT_CONTROL = 0xFFE3
		K_RIGHT_CONTROL = 0xFFE4
		K_ALT = 0xFFE9
		K_SPACE = 0x20
		K_ESC = 0xFF1B
		K_BACKSPACE = 0xFF08
		K_INSERT = 0xFF63
		K_DELETE = 0xFFFF
		K_HOME = 0xFF50
		K_END = 0xFF57
		K_PAGEUP = 0xFF55
		K_PAGEDOWN = 0xFF56
		K_LEFT = 0xFF51
		K_UP = 0xFF52
		K_RIGHT = 0xFF53
		K_DOWN = 0xFF54
		K_CAPSLOCK = 0xFFE5
		K_NUMLOCK = 0xFF7F
		K_SCROLLLOCK = 0xFF14
		K_LEFT_WIN = 0xFFEB
		K_RIGHT_WIN = 0xFFEC
		K_APPS = 0xFF67
		K_PAUSE = 0xFF13
		MOUSE_LEFT_BUTTON = 0x01
		MOUSE_RIGHT_BUTTON = 0x02
	}

}
