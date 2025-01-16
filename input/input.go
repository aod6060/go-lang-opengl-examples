package input

import sdl "github.com/veandco/go-sdl2/sdl"
import lm "derf.space/gldemo/linearmath"

// InputState
type InputState int32
const (
	IS_RELEASED InputState = iota
	IS_PRESSED_ONCE
	IS_PRESSED
	IS_RELEASED_ONCE
)

// Keyboard
type Keyboard int32
const (
	KEYS_UNKNOWN Keyboard = 0
	KEYS_A Keyboard = 4
	KEYS_B Keyboard = 5
	KEYS_C Keyboard = 6
	KEYS_D Keyboard = 7
	KEYS_E Keyboard = 8
	KEYS_F Keyboard = 9
	KEYS_G Keyboard = 10
	KEYS_H Keyboard = 11
	KEYS_I Keyboard = 12
	KEYS_J Keyboard = 13
	KEYS_K Keyboard = 14
	KEYS_L Keyboard = 15
	KEYS_M Keyboard = 16
	KEYS_N Keyboard = 17
	KEYS_O Keyboard = 18
	KEYS_P Keyboard = 19
	KEYS_Q Keyboard = 20
	KEYS_R Keyboard = 21
	KEYS_S Keyboard = 22
	KEYS_T Keyboard = 23
	KEYS_U Keyboard = 24
	KEYS_V Keyboard = 25
	KEYS_W Keyboard = 26
	KEYS_X Keyboard = 27
	KEYS_Y Keyboard = 28
	KEYS_Z Keyboard = 29
	KEYS_1 Keyboard = 30
	KEYS_2 Keyboard = 31
	KEYS_3 Keyboard = 32
	KEYS_4 Keyboard = 33
	KEYS_5 Keyboard = 34
	KEYS_6 Keyboard = 35
	KEYS_7 Keyboard = 36
	KEYS_8 Keyboard = 37
	KEYS_9 Keyboard = 38
	KEYS_0 Keyboard = 39
	KEYS_RETURN Keyboard = 40
	KEYS_ESCAPE Keyboard = 41
	KEYS_BACKSPACE Keyboard = 42
	KEYS_TAB Keyboard = 43
	KEYS_SPACE Keyboard = 44
	KEYS_MINUS Keyboard = 45
	KEYS_EQUALS Keyboard = 46
	KEYS_LEFTBRACKET Keyboard = 47
	KEYS_RIGHTBRACKET Keyboard = 48
	KEYS_BACKSLASH Keyboard = 49
	KEYS_NONUSHASH Keyboard = 50
	KEYS_SEMICOLON Keyboard = 51
	KEYS_APOSTROPHE Keyboard = 52
	KEYS_GRAVE Keyboard = 53
	KEYS_COMMA Keyboard = 54
	KEYS_PERIOD Keyboard = 55
	KEYS_SLASH Keyboard = 56
	KEYS_CAPSLOCK Keyboard = 57
	KEYS_F1 Keyboard = 58
	KEYS_F2 Keyboard = 59
	KEYS_F3 Keyboard = 60
	KEYS_F4 Keyboard = 61
	KEYS_F5 Keyboard = 62
	KEYS_F6 Keyboard = 63
	KEYS_F7 Keyboard = 64
	KEYS_F8 Keyboard = 65
	KEYS_F9 Keyboard = 66
	KEYS_F10 Keyboard = 67
	KEYS_F11 Keyboard = 68
	KEYS_F12 Keyboard = 69
	KEYS_PRINTSCREEN Keyboard = 70
	KEYS_SCROLLLOCK Keyboard = 71
	KEYS_PAUSE Keyboard = 72
	KEYS_INSERT Keyboard = 73
	KEYS_HOME Keyboard = 74
	KEYS_PAGEUP Keyboard = 75
	KEYS_DELETE Keyboard = 76
	KEYS_END Keyboard = 77
	KEYS_PAGEDOWN Keyboard = 78
	KEYS_RIGHT Keyboard = 79
	KEYS_LEFT Keyboard = 80
	KEYS_DOWN Keyboard = 81
	KEYS_UP Keyboard = 82
	KEYS_NUMLOCKCLEAR Keyboard = 83
	KEYS_KP_DIVIDE Keyboard = 84
	KEYS_KP_MULTIPLY Keyboard = 85
	KEYS_KP_MINUS Keyboard = 86
	KEYS_KP_PLUS Keyboard = 87
	KEYS_KP_ENTER Keyboard = 88
	KEYS_KP_1 Keyboard = 89
	KEYS_KP_2 Keyboard = 90
	KEYS_KP_3 Keyboard = 91
	KEYS_KP_4 Keyboard = 92
	KEYS_KP_5 Keyboard = 93
	KEYS_KP_6 Keyboard = 94
	KEYS_KP_7 Keyboard = 95
	KEYS_KP_8 Keyboard = 96
	KEYS_KP_9 Keyboard = 97
	KEYS_KP_0 Keyboard = 98
	KEYS_KP_PERIOD Keyboard = 99
	KEYS_NONUSBACKSLASH Keyboard = 100
	KEYS_APPLICATION Keyboard = 101
	KEYS_POWER Keyboard = 102
	KEYS_KP_EQUALS Keyboard = 103
	KEYS_F13 Keyboard = 104
	KEYS_F14 Keyboard = 105
	KEYS_F15 Keyboard = 106
	KEYS_F16 Keyboard = 107
	KEYS_F17 Keyboard = 108
	KEYS_F18 Keyboard = 109
	KEYS_F19 Keyboard = 110
	KEYS_F20 Keyboard = 111
	KEYS_F21 Keyboard = 112
	KEYS_F22 Keyboard = 113
	KEYS_F23 Keyboard = 114
	KEYS_F24 Keyboard = 115
	KEYS_EXECUTE Keyboard = 116
	KEYS_HELP Keyboard = 117
	KEYS_MENU Keyboard = 118
	KEYS_SELECT Keyboard = 119
	KEYS_STOP Keyboard = 120
	KEYS_AGAIN Keyboard = 121
	KEYS_UNDO Keyboard = 122
	KEYS_CUT Keyboard = 123
	KEYS_COPY Keyboard = 124
	KEYS_PASTE Keyboard = 125
	KEYS_FIND Keyboard = 126
	KEYS_MUTE Keyboard = 127
	KEYS_VOLUMEUP Keyboard = 128
	KEYS_VOLUMEDOWN Keyboard = 129
	KEYS_KP_COMMA Keyboard = 133
	KEYS_KP_EQUALSAS400 Keyboard = 134
	KEYS_INTERNATIONAL1 Keyboard = 135
	KEYS_INTERNATIONAL2 Keyboard = 136
	KEYS_INTERNATIONAL3 Keyboard = 137
	KEYS_INTERNATIONAL4 Keyboard = 138
	KEYS_INTERNATIONAL5 Keyboard = 139
	KEYS_INTERNATIONAL6 Keyboard = 140
	KEYS_INTERNATIONAL7 Keyboard = 141
	KEYS_INTERNATIONAL8 Keyboard = 142
	KEYS_INTERNATIONAL9 Keyboard = 143
	KEYS_LANG1 Keyboard = 144
	KEYS_LANG2 Keyboard = 145
	KEYS_LANG3 Keyboard = 146
	KEYS_LANG4 Keyboard = 147
	KEYS_LANG5 Keyboard = 148
	KEYS_LANG6 Keyboard = 149
	KEYS_LANG7 Keyboard = 150
	KEYS_LANG8 Keyboard = 151
	KEYS_LANG9 Keyboard = 152
	KEYS_ALTERASE Keyboard = 153
	KEYS_SYSREQ Keyboard = 154
	KEYS_CANCEL Keyboard = 155
	KEYS_CLEAR Keyboard = 156
	KEYS_PRIOR Keyboard = 157
	KEYS_RETURN2 Keyboard = 158
	KEYS_SEPARATOR Keyboard = 159
	KEYS_OUT Keyboard = 160
	KEYS_OPER Keyboard = 161
	KEYS_CLEARAGAIN Keyboard = 162
	KEYS_CRSEL Keyboard = 163
	KEYS_EXSEL Keyboard = 164
	KEYS_KP_00 Keyboard = 176
	KEYS_KP_000 Keyboard = 177
	KEYS_THOUSANDSSEPARATOR Keyboard = 178
	KEYS_DECIMALSEPARATOR Keyboard = 179
	KEYS_CURRENCYUNIT Keyboard = 180
	KEYS_CURRENCYSUBUNIT Keyboard = 181
	KEYS_KP_LEFTPAREN Keyboard = 182
	KEYS_KP_RIGHTPAREN Keyboard = 183
	KEYS_KP_LEFTBRACE Keyboard = 184
	KEYS_KP_RIGHTBRACE Keyboard = 185
	KEYS_KP_TAB Keyboard = 186
	KEYS_KP_BACKSPACE Keyboard = 187
	KEYS_KP_A Keyboard = 188
	KEYS_KP_B Keyboard = 189
	KEYS_KP_C Keyboard = 190
	KEYS_KP_D Keyboard = 191
	KEYS_KP_E Keyboard = 192
	KEYS_KP_F Keyboard = 193
	KEYS_KP_XOR Keyboard = 194
	KEYS_KP_POWER Keyboard = 195
	KEYS_KP_PERCENT Keyboard = 196
	KEYS_KP_LESS Keyboard = 197
	KEYS_KP_GREATER Keyboard = 198
	KEYS_KP_AMPERSAND Keyboard = 199
	KEYS_KP_DBLAMPERSAND Keyboard = 200
	KEYS_KP_VERTICALBAR Keyboard = 201
	KEYS_KP_DBLVERTICALBAR Keyboard = 202
	KEYS_KP_COLON Keyboard = 203
	KEYS_KP_HASH Keyboard = 204
	KEYS_KP_SPACE Keyboard = 205
	KEYS_KP_AT Keyboard = 206
	KEYS_KP_EXCLAM Keyboard = 207
	KEYS_KP_MEMSTORE Keyboard = 208
	KEYS_KP_MEMRECALL Keyboard = 209
	KEYS_KP_MEMCLEAR Keyboard = 210
	KEYS_KP_MEMADD Keyboard = 211
	KEYS_KP_MEMSUBTRACT Keyboard = 212
	KEYS_KP_MEMMULTIPLY Keyboard = 213
	KEYS_KP_MEMDIVIDE Keyboard = 214
	KEYS_KP_PLUSMINUS Keyboard = 215
	KEYS_KP_CLEAR Keyboard = 216
	KEYS_KP_CLEARENTRY Keyboard = 217
	KEYS_KP_BINARY Keyboard = 218
	KEYS_KP_OCTAL Keyboard = 219
	KEYS_KP_DECIMAL Keyboard = 220
	KEYS_KP_HEXADECIMAL Keyboard = 221
	KEYS_LCTRL Keyboard = 224
	KEYS_LSHIFT Keyboard = 225
	KEYS_LALT Keyboard = 226
	KEYS_LGUI Keyboard = 227
	KEYS_RCTRL Keyboard = 228
	KEYS_RSHIFT Keyboard = 229
	KEYS_RALT Keyboard = 230
	KEYS_RGUI Keyboard = 231
	KEYS_MODE Keyboard = 257
	KEYS_AUDIONEXT Keyboard = 258
	KEYS_AUDIOPREV Keyboard = 259
	KEYS_AUDIOSTOP Keyboard = 260
	KEYS_AUDIOPLAY Keyboard = 261
	KEYS_AUDIOMUTE Keyboard = 262
	KEYS_MEDIASELECT Keyboard = 263
	KEYS_WWW Keyboard = 264
	KEYS_MAIL Keyboard = 265
	KEYS_CALCULATOR Keyboard = 266
	KEYS_COMPUTER Keyboard = 267
	KEYS_AC_SEARCH Keyboard = 268
	KEYS_AC_HOME Keyboard = 269
	KEYS_AC_BACK Keyboard = 270
	KEYS_AC_FORWARD Keyboard = 271
	KEYS_AC_STOP Keyboard = 272
	KEYS_AC_REFRESH Keyboard = 273
	KEYS_AC_BOOKMARKS Keyboard = 274
	KEYS_BRIGHTNESSDOWN Keyboard = 275
	KEYS_BRIGHTNESSUP Keyboard = 276
	KEYS_DISPLAYSWITCH Keyboard = 277
	KEYS_KBDILLUMTOGGLE Keyboard = 278
	KEYS_KBDILLUMDOWN Keyboard = 279
	KEYS_KBDILLUMUP Keyboard = 280
	KEYS_EJECT Keyboard = 281
	KEYS_SLEEP Keyboard = 282
	KEYS_APP1 Keyboard = 283
	KEYS_APP2 Keyboard = 284
	KEYS_AUDIOREWIND Keyboard = 285
	KEYS_AUDIOFASTFORWARD Keyboard = 286
	KEYS_SOFTLEFT Keyboard = 287
	KEYS_SOFTRIGHT Keyboard = 288
	KEYS_ENDCALL Keyboard = 290
	KEYS_MAX_SIZE Keyboard = 512
)

type MouseButtons int32
const (
	MBS_LEFT MouseButtons = iota
	MBS_CENTER
	MBS_RIGHT
	MBS_MAX_SIZE
)


var keys []InputState

var mouseButtons []InputState
var position *lm.Vec2 = lm.CreateVec2(0.0, 0.0)
var relative *lm.Vec2 = lm.CreateVec2(0.0, 0.0)
var wheel *lm.Vec2    = lm.CreateVec2(0.0, 0.0)

var time float32 = 1.0
var maxTime float32 = 1.0

func Init() {
	keys = make([]InputState, KEYS_MAX_SIZE)

	for i := 0; i < int(KEYS_MAX_SIZE); i++ {
		keys[i] = IS_RELEASED
	}

	mouseButtons = make([]InputState, MBS_MAX_SIZE)

	for i := 0; i < int(MBS_MAX_SIZE); i++ {
		mouseButtons[i] = IS_RELEASED
	}
}

func HandleEvent(e sdl.Event) {
	if e.GetType() == sdl.KEYDOWN {
		var temp *sdl.KeyboardEvent = e.(*sdl.KeyboardEvent)
		if keys[temp.Keysym.Scancode] == IS_RELEASED {
			keys[temp.Keysym.Scancode] = IS_PRESSED_ONCE
		}
	} else if e.GetType() == sdl.KEYUP {
		var temp *sdl.KeyboardEvent = e.(*sdl.KeyboardEvent)
		if keys[temp.Keysym.Scancode] == IS_PRESSED {
			keys[temp.Keysym.Scancode] = IS_RELEASED_ONCE
		}
	} else if e.GetType() == sdl.MOUSEBUTTONDOWN {
		var temp *sdl.MouseButtonEvent = e.(*sdl.MouseButtonEvent)
		if mouseButtons[temp.Button - 1] == IS_RELEASED {
			mouseButtons[temp.Button - 1] = IS_PRESSED_ONCE
		}
	} else if e.GetType() == sdl.MOUSEBUTTONUP {
		var temp *sdl.MouseButtonEvent = e.(*sdl.MouseButtonEvent)
		if mouseButtons[temp.Button - 1] == IS_PRESSED {
			mouseButtons[temp.Button - 1] = IS_RELEASED_ONCE
		}
	} else if e.GetType() == sdl.MOUSEMOTION {
		var temp *sdl.MouseMotionEvent = e.(*sdl.MouseMotionEvent)
		position.X = float32(temp.X)
		position.Y = float32(temp.Y)
		relative.X = float32(temp.XRel)
		relative.Y = float32(temp.YRel)
		time = 0.0
	} else if e.GetType() == sdl.MOUSEWHEEL {
		var temp *sdl.MouseWheelEvent = e.(*sdl.MouseWheelEvent)
		wheel.X = float32(temp.X)
		wheel.Y = float32(temp.Y)
	}
}

func Update(delta float32) {
	// Keyboard
	for i:= 0; i < int(KEYS_MAX_SIZE); i++ {
		if keys[i] == IS_PRESSED_ONCE {
			keys[i] = IS_PRESSED
		}

		if keys[i] == IS_RELEASED_ONCE {
			keys[i] = IS_RELEASED
		}
	}

	// MouseButton
	for i:= 0; i < int(MBS_MAX_SIZE); i++ {
		if mouseButtons[i] == IS_PRESSED_ONCE {
			mouseButtons[i] = IS_PRESSED
		}

		if mouseButtons[i] == IS_RELEASED_ONCE {
			mouseButtons[i] = IS_RELEASED
		}
	}

	if(time < maxTime) {
		time += delta
	}

	var t = time / maxTime

	relative.X = (0.0 - relative.X) * (t * t) + relative.X
	relative.Y = (0.0 - relative.Y) * (t * t) + relative.Y

	wheel.X = 0.0
	wheel.Y = 0.0
}

func Release() {
	keys = keys[:0]
	mouseButtons = mouseButtons[:0]
}

func IsKeyReleased(key Keyboard) bool {
	return keys[key] == IS_RELEASED || keys[key] == IS_RELEASED_ONCE
}

func IsKeyPressedOnce(key Keyboard) bool {
	return keys[key] == IS_PRESSED_ONCE
}

func IsKeyPressed(key Keyboard) bool {
	return keys[key] == IS_PRESSED || keys[key] == IS_PRESSED_ONCE
}

func IsKeyReleasedOnce(key Keyboard) bool {
	return keys[key] == IS_RELEASED_ONCE
}

func GetKeyReleasedValue(key Keyboard) float32 {
	var value float32 = 0.0
	if IsKeyReleased(key) {
		value = 1.0
	}
	return value
}

func GetKeyPressedOnceValue(key Keyboard) float32 {
	var value float32 = 0.0
	if IsKeyPressedOnce(key) {
		value = 1.0
	}
	return value
}

func GetKeyPressedValue(key Keyboard) float32 {
	var value float32 = 0.0
	if IsKeyPressed(key) {
		value = 1.0
	}
	return value
}

func GetKeyReleasedOnceValue(key Keyboard) float32 {
	var value float32 = 0.0
	if IsKeyReleasedOnce(key) {
		value = 1.0
	}
	return value
}

func GetKeyReleasedAxis(negative Keyboard, positive Keyboard) float32 {
	return GetKeyReleasedValue(positive) - GetKeyReleasedValue(negative)
}

func GetKeyPressedOnceAxis(negative Keyboard, positive Keyboard) float32 {
	return GetKeyPressedOnceValue(positive) - GetKeyPressedOnceValue(negative)
}

func GetKeyPressedAxis(negative Keyboard, positive Keyboard) float32 {
	return GetKeyPressedValue(positive) - GetKeyPressedValue(negative)
}

func GetKeyReleasedOnceAxis(negative Keyboard, positive Keyboard) float32 {
	return GetKeyReleasedOnceValue(positive) - GetKeyReleasedOnceValue(negative)
}

func IsMouseButtonReleased(b MouseButtons) bool {
	return mouseButtons[b] == IS_RELEASED || mouseButtons[b] == IS_RELEASED_ONCE
}

func IsMouseButtonPressedOnce(b MouseButtons) bool {
	return mouseButtons[b] == IS_PRESSED_ONCE
}

func IsMouseButtonPressed(b MouseButtons) bool {
	return mouseButtons[b] == IS_PRESSED || mouseButtons[b] == IS_PRESSED_ONCE
}

func IsMouseButtonReleasedOnce(b MouseButtons) bool {
	return mouseButtons[b] == IS_RELEASED_ONCE
}

func GetMouseButtonReleasedValue(b MouseButtons) float32 {
	var value float32 = 0.0
	if IsMouseButtonReleased(b) {
		value = 1.0
	}
	return value
}

func GetMouseButtonPressedOnceValue(b MouseButtons) float32 {
	var value float32 = 0.0
	if IsMouseButtonPressedOnce(b) {
		value = 1.0
	}
	return value
}

func GetMouseButtonPressedValue(b MouseButtons) float32 {
	var value float32 = 0.0
	if IsMouseButtonPressed(b) {
		value = 1.0
	}
	return value
}

func GetMouseButtonReleasedOnceValue(b MouseButtons) float32 {
	var value float32 = 0.0
	if IsMouseButtonReleasedOnce(b) {
		value = 1.0
	}
	return value
}

func GetMouseButtonReleasedAxis(negative MouseButtons, positive MouseButtons) float32 {
	return GetMouseButtonReleasedValue(positive) - GetMouseButtonReleasedValue(negative)
}

func GetMouseButtonPressedOnceAxis(negative MouseButtons, positive MouseButtons) float32 {
	return GetMouseButtonPressedOnceValue(positive) - GetMouseButtonPressedOnceValue(negative)
}

func GetMouseButtonPressedAxis(negative MouseButtons, positive MouseButtons) float32 {
	return GetMouseButtonPressedValue(positive) - GetMouseButtonPressedValue(negative)
}

func GetMouseButtonReleasedOnceAxis(negative MouseButtons, positive MouseButtons) float32 {
	return GetMouseButtonReleasedOnceValue(positive) - GetMouseButtonReleasedOnceValue(negative)
}

func GetPosition() *lm.Vec2 {
	return position
}

func GetRelative() *lm.Vec2 {
	return relative
}

func GetMouseWheel() *lm.Vec2 {
	return wheel
}

func SetGrab(grab bool) {
	sdl.SetRelativeMouseMode(grab)
}

func GetGrab() bool {
	return sdl.GetRelativeMouseMode()
}

func ToggleGrab() {
	SetGrab(!GetGrab())
}