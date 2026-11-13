package cio

import b "github.com/vault-thirteen/auxie/os/win32/base"

// KeyCode is a virtual key code.
// This functionality is guaranteed to work only on Windows operating systems.
type KeyCode byte

// Virtual Key codes.
const (
	VK_LBUTTON                         = KeyCode(b.VK_LBUTTON)
	VK_RBUTTON                         = KeyCode(b.VK_RBUTTON)
	VK_CANCEL                          = KeyCode(b.VK_CANCEL)
	VK_MBUTTON                         = KeyCode(b.VK_MBUTTON)
	VK_XBUTTON1                        = KeyCode(b.VK_XBUTTON1)
	VK_XBUTTON2                        = KeyCode(b.VK_XBUTTON2)
	VK_BACK                            = KeyCode(b.VK_BACK)
	VK_TAB                             = KeyCode(b.VK_TAB)
	VK_CLEAR                           = KeyCode(b.VK_CLEAR)
	VK_RETURN                          = KeyCode(b.VK_RETURN)
	VK_SHIFT                           = KeyCode(b.VK_SHIFT)
	VK_CONTROL                         = KeyCode(b.VK_CONTROL)
	VK_MENU                            = KeyCode(b.VK_MENU)
	VK_PAUSE                           = KeyCode(b.VK_PAUSE)
	VK_CAPITAL                         = KeyCode(b.VK_CAPITAL)
	VK_KANA                            = KeyCode(b.VK_KANA)
	VK_HANGUL                          = KeyCode(b.VK_HANGUL)
	VK_IME_ON                          = KeyCode(b.VK_IME_ON)
	VK_JUNJA                           = KeyCode(b.VK_JUNJA)
	VK_FINAL                           = KeyCode(b.VK_FINAL)
	VK_HANJA                           = KeyCode(b.VK_HANJA)
	VK_KANJI                           = KeyCode(b.VK_KANJI)
	VK_IME_OFF                         = KeyCode(b.VK_IME_OFF)
	VK_ESCAPE                          = KeyCode(b.VK_ESCAPE)
	VK_CONVERT                         = KeyCode(b.VK_CONVERT)
	VK_NONCONVERT                      = KeyCode(b.VK_NONCONVERT)
	VK_ACCEPT                          = KeyCode(b.VK_ACCEPT)
	VK_MODECHANGE                      = KeyCode(b.VK_MODECHANGE)
	VK_SPACE                           = KeyCode(b.VK_SPACE)
	VK_PRIOR                           = KeyCode(b.VK_PRIOR)
	VK_NEXT                            = KeyCode(b.VK_NEXT)
	VK_END                             = KeyCode(b.VK_END)
	VK_HOME                            = KeyCode(b.VK_HOME)
	VK_LEFT                            = KeyCode(b.VK_LEFT)
	VK_UP                              = KeyCode(b.VK_UP)
	VK_RIGHT                           = KeyCode(b.VK_RIGHT)
	VK_DOWN                            = KeyCode(b.VK_DOWN)
	VK_SELECT                          = KeyCode(b.VK_SELECT)
	VK_PRINT                           = KeyCode(b.VK_PRINT)
	VK_EXECUTE                         = KeyCode(b.VK_EXECUTE)
	VK_SNAPSHOT                        = KeyCode(b.VK_SNAPSHOT)
	VK_INSERT                          = KeyCode(b.VK_INSERT)
	VK_DELETE                          = KeyCode(b.VK_DELETE)
	VK_HELP                            = KeyCode(b.VK_HELP)
	VK_0                               = KeyCode(b.VK_0)
	VK_1                               = KeyCode(b.VK_1)
	VK_2                               = KeyCode(b.VK_2)
	VK_3                               = KeyCode(b.VK_3)
	VK_4                               = KeyCode(b.VK_4)
	VK_5                               = KeyCode(b.VK_5)
	VK_6                               = KeyCode(b.VK_6)
	VK_7                               = KeyCode(b.VK_7)
	VK_8                               = KeyCode(b.VK_8)
	VK_9                               = KeyCode(b.VK_9)
	VK_A                               = KeyCode(b.VK_A)
	VK_B                               = KeyCode(b.VK_B)
	VK_C                               = KeyCode(b.VK_C)
	VK_D                               = KeyCode(b.VK_D)
	VK_E                               = KeyCode(b.VK_E)
	VK_F                               = KeyCode(b.VK_F)
	VK_G                               = KeyCode(b.VK_G)
	VK_H                               = KeyCode(b.VK_H)
	VK_I                               = KeyCode(b.VK_I)
	VK_J                               = KeyCode(b.VK_J)
	VK_K                               = KeyCode(b.VK_K)
	VK_L                               = KeyCode(b.VK_L)
	VK_M                               = KeyCode(b.VK_M)
	VK_N                               = KeyCode(b.VK_N)
	VK_O                               = KeyCode(b.VK_O)
	VK_P                               = KeyCode(b.VK_P)
	VK_Q                               = KeyCode(b.VK_Q)
	VK_R                               = KeyCode(b.VK_R)
	VK_S                               = KeyCode(b.VK_S)
	VK_T                               = KeyCode(b.VK_T)
	VK_U                               = KeyCode(b.VK_U)
	VK_V                               = KeyCode(b.VK_V)
	VK_W                               = KeyCode(b.VK_W)
	VK_X                               = KeyCode(b.VK_X)
	VK_Y                               = KeyCode(b.VK_Y)
	VK_Z                               = KeyCode(b.VK_Z)
	VK_LWIN                            = KeyCode(b.VK_LWIN)
	VK_RWIN                            = KeyCode(b.VK_RWIN)
	VK_APPS                            = KeyCode(b.VK_APPS)
	VK_SLEEP                           = KeyCode(b.VK_SLEEP)
	VK_NUMPAD0                         = KeyCode(b.VK_NUMPAD0)
	VK_NUMPAD1                         = KeyCode(b.VK_NUMPAD1)
	VK_NUMPAD2                         = KeyCode(b.VK_NUMPAD2)
	VK_NUMPAD3                         = KeyCode(b.VK_NUMPAD3)
	VK_NUMPAD4                         = KeyCode(b.VK_NUMPAD4)
	VK_NUMPAD5                         = KeyCode(b.VK_NUMPAD5)
	VK_NUMPAD6                         = KeyCode(b.VK_NUMPAD6)
	VK_NUMPAD7                         = KeyCode(b.VK_NUMPAD7)
	VK_NUMPAD8                         = KeyCode(b.VK_NUMPAD8)
	VK_NUMPAD9                         = KeyCode(b.VK_NUMPAD9)
	VK_MULTIPLY                        = KeyCode(b.VK_MULTIPLY)
	VK_ADD                             = KeyCode(b.VK_ADD)
	VK_SEPARATOR                       = KeyCode(b.VK_SEPARATOR)
	VK_SUBTRACT                        = KeyCode(b.VK_SUBTRACT)
	VK_DECIMAL                         = KeyCode(b.VK_DECIMAL)
	VK_DIVIDE                          = KeyCode(b.VK_DIVIDE)
	VK_F1                              = KeyCode(b.VK_F1)
	VK_F2                              = KeyCode(b.VK_F2)
	VK_F3                              = KeyCode(b.VK_F3)
	VK_F4                              = KeyCode(b.VK_F4)
	VK_F5                              = KeyCode(b.VK_F5)
	VK_F6                              = KeyCode(b.VK_F6)
	VK_F7                              = KeyCode(b.VK_F7)
	VK_F8                              = KeyCode(b.VK_F8)
	VK_F9                              = KeyCode(b.VK_F9)
	VK_F10                             = KeyCode(b.VK_F10)
	VK_F11                             = KeyCode(b.VK_F11)
	VK_F12                             = KeyCode(b.VK_F12)
	VK_F13                             = KeyCode(b.VK_F13)
	VK_F14                             = KeyCode(b.VK_F14)
	VK_F15                             = KeyCode(b.VK_F15)
	VK_F16                             = KeyCode(b.VK_F16)
	VK_F17                             = KeyCode(b.VK_F17)
	VK_F18                             = KeyCode(b.VK_F18)
	VK_F19                             = KeyCode(b.VK_F19)
	VK_F20                             = KeyCode(b.VK_F20)
	VK_F21                             = KeyCode(b.VK_F21)
	VK_F22                             = KeyCode(b.VK_F22)
	VK_F23                             = KeyCode(b.VK_F23)
	VK_F24                             = KeyCode(b.VK_F24)
	VK_NUMLOCK                         = KeyCode(b.VK_NUMLOCK)
	VK_SCROLL                          = KeyCode(b.VK_SCROLL)
	VK_LSHIFT                          = KeyCode(b.VK_LSHIFT)
	VK_RSHIFT                          = KeyCode(b.VK_RSHIFT)
	VK_LCONTROL                        = KeyCode(b.VK_LCONTROL)
	VK_RCONTROL                        = KeyCode(b.VK_RCONTROL)
	VK_LMENU                           = KeyCode(b.VK_LMENU)
	VK_RMENU                           = KeyCode(b.VK_RMENU)
	VK_BROWSER_BACK                    = KeyCode(b.VK_BROWSER_BACK)
	VK_BROWSER_FORWARD                 = KeyCode(b.VK_BROWSER_FORWARD)
	VK_BROWSER_REFRESH                 = KeyCode(b.VK_BROWSER_REFRESH)
	VK_BROWSER_STOP                    = KeyCode(b.VK_BROWSER_STOP)
	VK_BROWSER_SEARCH                  = KeyCode(b.VK_BROWSER_SEARCH)
	VK_BROWSER_FAVORITES               = KeyCode(b.VK_BROWSER_FAVORITES)
	VK_BROWSER_HOME                    = KeyCode(b.VK_BROWSER_HOME)
	VK_VOLUME_MUTE                     = KeyCode(b.VK_VOLUME_MUTE)
	VK_VOLUME_DOWN                     = KeyCode(b.VK_VOLUME_DOWN)
	VK_VOLUME_UP                       = KeyCode(b.VK_VOLUME_UP)
	VK_MEDIA_NEXT_TRACK                = KeyCode(b.VK_MEDIA_NEXT_TRACK)
	VK_MEDIA_PREV_TRACK                = KeyCode(b.VK_MEDIA_PREV_TRACK)
	VK_MEDIA_STOP                      = KeyCode(b.VK_MEDIA_STOP)
	VK_MEDIA_PLAY_PAUSE                = KeyCode(b.VK_MEDIA_PLAY_PAUSE)
	VK_LAUNCH_MAIL                     = KeyCode(b.VK_LAUNCH_MAIL)
	VK_LAUNCH_MEDIA_SELECT             = KeyCode(b.VK_LAUNCH_MEDIA_SELECT)
	VK_LAUNCH_APP1                     = KeyCode(b.VK_LAUNCH_APP1)
	VK_LAUNCH_APP2                     = KeyCode(b.VK_LAUNCH_APP2)
	VK_OEM_1                           = KeyCode(b.VK_OEM_1)
	VK_OEM_PLUS                        = KeyCode(b.VK_OEM_PLUS)
	VK_OEM_COMMA                       = KeyCode(b.VK_OEM_COMMA)
	VK_OEM_MINUS                       = KeyCode(b.VK_OEM_MINUS)
	VK_OEM_PERIOD                      = KeyCode(b.VK_OEM_PERIOD)
	VK_OEM_2                           = KeyCode(b.VK_OEM_2)
	VK_OEM_3                           = KeyCode(b.VK_OEM_3)
	VK_GAMEPAD_A                       = KeyCode(b.VK_GAMEPAD_A)
	VK_GAMEPAD_B                       = KeyCode(b.VK_GAMEPAD_B)
	VK_GAMEPAD_X                       = KeyCode(b.VK_GAMEPAD_X)
	VK_GAMEPAD_Y                       = KeyCode(b.VK_GAMEPAD_Y)
	VK_GAMEPAD_RIGHT_SHOULDER          = KeyCode(b.VK_GAMEPAD_RIGHT_SHOULDER)
	VK_GAMEPAD_LEFT_SHOULDER           = KeyCode(b.VK_GAMEPAD_LEFT_SHOULDER)
	VK_GAMEPAD_LEFT_TRIGGER            = KeyCode(b.VK_GAMEPAD_LEFT_TRIGGER)
	VK_GAMEPAD_RIGHT_TRIGGER           = KeyCode(b.VK_GAMEPAD_RIGHT_TRIGGER)
	VK_GAMEPAD_DPAD_UP                 = KeyCode(b.VK_GAMEPAD_DPAD_UP)
	VK_GAMEPAD_DPAD_DOWN               = KeyCode(b.VK_GAMEPAD_DPAD_DOWN)
	VK_GAMEPAD_DPAD_LEFT               = KeyCode(b.VK_GAMEPAD_DPAD_LEFT)
	VK_GAMEPAD_DPAD_RIGHT              = KeyCode(b.VK_GAMEPAD_DPAD_RIGHT)
	VK_GAMEPAD_MENU                    = KeyCode(b.VK_GAMEPAD_MENU)
	VK_GAMEPAD_VIEW                    = KeyCode(b.VK_GAMEPAD_VIEW)
	VK_GAMEPAD_LEFT_THUMBSTICK_BUTTON  = KeyCode(b.VK_GAMEPAD_LEFT_THUMBSTICK_BUTTON)
	VK_GAMEPAD_RIGHT_THUMBSTICK_BUTTON = KeyCode(b.VK_GAMEPAD_RIGHT_THUMBSTICK_BUTTON)
	VK_GAMEPAD_LEFT_THUMBSTICK_UP      = KeyCode(b.VK_GAMEPAD_LEFT_THUMBSTICK_UP)
	VK_GAMEPAD_LEFT_THUMBSTICK_DOWN    = KeyCode(b.VK_GAMEPAD_LEFT_THUMBSTICK_DOWN)
	VK_GAMEPAD_LEFT_THUMBSTICK_RIGHT   = KeyCode(b.VK_GAMEPAD_LEFT_THUMBSTICK_RIGHT)
	VK_GAMEPAD_LEFT_THUMBSTICK_LEFT    = KeyCode(b.VK_GAMEPAD_LEFT_THUMBSTICK_LEFT)
	VK_GAMEPAD_RIGHT_THUMBSTICK_UP     = KeyCode(b.VK_GAMEPAD_RIGHT_THUMBSTICK_UP)
	VK_GAMEPAD_RIGHT_THUMBSTICK_DOWN   = KeyCode(b.VK_GAMEPAD_RIGHT_THUMBSTICK_DOWN)
	VK_GAMEPAD_RIGHT_THUMBSTICK_RIGHT  = KeyCode(b.VK_GAMEPAD_RIGHT_THUMBSTICK_RIGHT)
	VK_GAMEPAD_RIGHT_THUMBSTICK_LEFT   = KeyCode(b.VK_GAMEPAD_RIGHT_THUMBSTICK_LEFT)
	VK_OEM_4                           = KeyCode(b.VK_OEM_4)
	VK_OEM_5                           = KeyCode(b.VK_OEM_5)
	VK_OEM_6                           = KeyCode(b.VK_OEM_6)
	VK_OEM_7                           = KeyCode(b.VK_OEM_7)
	VK_OEM_8                           = KeyCode(b.VK_OEM_8)
	VK_OEM_102                         = KeyCode(b.VK_OEM_102)
	VK_PROCESSKEY                      = KeyCode(b.VK_PROCESSKEY)
	VK_PACKET                          = KeyCode(b.VK_PACKET)
	VK_ATTN                            = KeyCode(b.VK_ATTN)
	VK_CRSEL                           = KeyCode(b.VK_CRSEL)
	VK_EXSEL                           = KeyCode(b.VK_EXSEL)
	VK_EREOF                           = KeyCode(b.VK_EREOF)
	VK_PLAY                            = KeyCode(b.VK_PLAY)
	VK_ZOOM                            = KeyCode(b.VK_ZOOM)
	VK_NONAME                          = KeyCode(b.VK_NONAME)
	VK_PA1                             = KeyCode(b.VK_PA1)
	VK_OEM_CLEAR                       = KeyCode(b.VK_OEM_CLEAR)
)

var KeyCodeNames = map[KeyCode]string{
	VK_LBUTTON:                         "VK_LBUTTON",
	VK_RBUTTON:                         "VK_RBUTTON",
	VK_CANCEL:                          "VK_CANCEL",
	VK_MBUTTON:                         "VK_MBUTTON",
	VK_XBUTTON1:                        "VK_XBUTTON1",
	VK_XBUTTON2:                        "VK_XBUTTON2",
	VK_BACK:                            "VK_BACK",
	VK_TAB:                             "VK_TAB",
	VK_CLEAR:                           "VK_CLEAR",
	VK_RETURN:                          "VK_RETURN",
	VK_SHIFT:                           "VK_SHIFT",
	VK_CONTROL:                         "VK_CONTROL",
	VK_MENU:                            "VK_MENU",
	VK_PAUSE:                           "VK_PAUSE",
	VK_CAPITAL:                         "VK_CAPITAL",
	VK_KANA:                            "VK_KANA",
	VK_IME_ON:                          "VK_IME_ON",
	VK_JUNJA:                           "VK_JUNJA",
	VK_FINAL:                           "VK_FINAL",
	VK_KANJI:                           "VK_KANJI",
	VK_IME_OFF:                         "VK_IME_OFF",
	VK_ESCAPE:                          "VK_ESCAPE",
	VK_CONVERT:                         "VK_CONVERT",
	VK_NONCONVERT:                      "VK_NONCONVERT",
	VK_ACCEPT:                          "VK_ACCEPT",
	VK_MODECHANGE:                      "VK_MODECHANGE",
	VK_SPACE:                           "VK_SPACE",
	VK_PRIOR:                           "VK_PRIOR",
	VK_NEXT:                            "VK_NEXT",
	VK_END:                             "VK_END",
	VK_HOME:                            "VK_HOME",
	VK_LEFT:                            "VK_LEFT",
	VK_UP:                              "VK_UP",
	VK_RIGHT:                           "VK_RIGHT",
	VK_DOWN:                            "VK_DOWN",
	VK_SELECT:                          "VK_SELECT",
	VK_PRINT:                           "VK_PRINT",
	VK_EXECUTE:                         "VK_EXECUTE",
	VK_SNAPSHOT:                        "VK_SNAPSHOT",
	VK_INSERT:                          "VK_INSERT",
	VK_DELETE:                          "VK_DELETE",
	VK_HELP:                            "VK_HELP",
	VK_0:                               "VK_0",
	VK_1:                               "VK_1",
	VK_2:                               "VK_2",
	VK_3:                               "VK_3",
	VK_4:                               "VK_4",
	VK_5:                               "VK_5",
	VK_6:                               "VK_6",
	VK_7:                               "VK_7",
	VK_8:                               "VK_8",
	VK_9:                               "VK_9",
	VK_A:                               "VK_A",
	VK_B:                               "VK_B",
	VK_C:                               "VK_C",
	VK_D:                               "VK_D",
	VK_E:                               "VK_E",
	VK_F:                               "VK_F",
	VK_G:                               "VK_G",
	VK_H:                               "VK_H",
	VK_I:                               "VK_I",
	VK_J:                               "VK_J",
	VK_K:                               "VK_K",
	VK_L:                               "VK_L",
	VK_M:                               "VK_M",
	VK_N:                               "VK_N",
	VK_O:                               "VK_O",
	VK_P:                               "VK_P",
	VK_Q:                               "VK_Q",
	VK_R:                               "VK_R",
	VK_S:                               "VK_S",
	VK_T:                               "VK_T",
	VK_U:                               "VK_U",
	VK_V:                               "VK_V",
	VK_W:                               "VK_W",
	VK_X:                               "VK_X",
	VK_Y:                               "VK_Y",
	VK_Z:                               "VK_Z",
	VK_LWIN:                            "VK_LWIN",
	VK_RWIN:                            "VK_RWIN",
	VK_APPS:                            "VK_APPS",
	VK_SLEEP:                           "VK_SLEEP",
	VK_NUMPAD0:                         "VK_NUMPAD0",
	VK_NUMPAD1:                         "VK_NUMPAD1",
	VK_NUMPAD2:                         "VK_NUMPAD2",
	VK_NUMPAD3:                         "VK_NUMPAD3",
	VK_NUMPAD4:                         "VK_NUMPAD4",
	VK_NUMPAD5:                         "VK_NUMPAD5",
	VK_NUMPAD6:                         "VK_NUMPAD6",
	VK_NUMPAD7:                         "VK_NUMPAD7",
	VK_NUMPAD8:                         "VK_NUMPAD8",
	VK_NUMPAD9:                         "VK_NUMPAD9",
	VK_MULTIPLY:                        "VK_MULTIPLY",
	VK_ADD:                             "VK_ADD",
	VK_SEPARATOR:                       "VK_SEPARATOR",
	VK_SUBTRACT:                        "VK_SUBTRACT",
	VK_DECIMAL:                         "VK_DECIMAL",
	VK_DIVIDE:                          "VK_DIVIDE",
	VK_F1:                              "VK_F1",
	VK_F2:                              "VK_F2",
	VK_F3:                              "VK_F3",
	VK_F4:                              "VK_F4",
	VK_F5:                              "VK_F5",
	VK_F6:                              "VK_F6",
	VK_F7:                              "VK_F7",
	VK_F8:                              "VK_F8",
	VK_F9:                              "VK_F9",
	VK_F10:                             "VK_F10",
	VK_F11:                             "VK_F11",
	VK_F12:                             "VK_F12",
	VK_F13:                             "VK_F13",
	VK_F14:                             "VK_F14",
	VK_F15:                             "VK_F15",
	VK_F16:                             "VK_F16",
	VK_F17:                             "VK_F17",
	VK_F18:                             "VK_F18",
	VK_F19:                             "VK_F19",
	VK_F20:                             "VK_F20",
	VK_F21:                             "VK_F21",
	VK_F22:                             "VK_F22",
	VK_F23:                             "VK_F23",
	VK_F24:                             "VK_F24",
	VK_NUMLOCK:                         "VK_NUMLOCK",
	VK_SCROLL:                          "VK_SCROLL",
	VK_LSHIFT:                          "VK_LSHIFT",
	VK_RSHIFT:                          "VK_RSHIFT",
	VK_LCONTROL:                        "VK_LCONTROL",
	VK_RCONTROL:                        "VK_RCONTROL",
	VK_LMENU:                           "VK_LMENU",
	VK_RMENU:                           "VK_RMENU",
	VK_BROWSER_BACK:                    "VK_BROWSER_BACK",
	VK_BROWSER_FORWARD:                 "VK_BROWSER_FORWARD",
	VK_BROWSER_REFRESH:                 "VK_BROWSER_REFRESH",
	VK_BROWSER_STOP:                    "VK_BROWSER_STOP",
	VK_BROWSER_SEARCH:                  "VK_BROWSER_SEARCH",
	VK_BROWSER_FAVORITES:               "VK_BROWSER_FAVORITES",
	VK_BROWSER_HOME:                    "VK_BROWSER_HOME",
	VK_VOLUME_MUTE:                     "VK_VOLUME_MUTE",
	VK_VOLUME_DOWN:                     "VK_VOLUME_DOWN",
	VK_VOLUME_UP:                       "VK_VOLUME_UP",
	VK_MEDIA_NEXT_TRACK:                "VK_MEDIA_NEXT_TRACK",
	VK_MEDIA_PREV_TRACK:                "VK_MEDIA_PREV_TRACK",
	VK_MEDIA_STOP:                      "VK_MEDIA_STOP",
	VK_MEDIA_PLAY_PAUSE:                "VK_MEDIA_PLAY_PAUSE",
	VK_LAUNCH_MAIL:                     "VK_LAUNCH_MAIL",
	VK_LAUNCH_MEDIA_SELECT:             "VK_LAUNCH_MEDIA_SELECT",
	VK_LAUNCH_APP1:                     "VK_LAUNCH_APP1",
	VK_LAUNCH_APP2:                     "VK_LAUNCH_APP2",
	VK_OEM_1:                           "VK_OEM_1",
	VK_OEM_PLUS:                        "VK_OEM_PLUS",
	VK_OEM_COMMA:                       "VK_OEM_COMMA",
	VK_OEM_MINUS:                       "VK_OEM_MINUS",
	VK_OEM_PERIOD:                      "VK_OEM_PERIOD",
	VK_OEM_2:                           "VK_OEM_2",
	VK_OEM_3:                           "VK_OEM_3",
	VK_GAMEPAD_A:                       "VK_GAMEPAD_A",
	VK_GAMEPAD_B:                       "VK_GAMEPAD_B",
	VK_GAMEPAD_X:                       "VK_GAMEPAD_X",
	VK_GAMEPAD_Y:                       "VK_GAMEPAD_Y",
	VK_GAMEPAD_RIGHT_SHOULDER:          "VK_GAMEPAD_RIGHT_SHOULDER",
	VK_GAMEPAD_LEFT_SHOULDER:           "VK_GAMEPAD_LEFT_SHOULDER",
	VK_GAMEPAD_LEFT_TRIGGER:            "VK_GAMEPAD_LEFT_TRIGGER",
	VK_GAMEPAD_RIGHT_TRIGGER:           "VK_GAMEPAD_RIGHT_TRIGGER",
	VK_GAMEPAD_DPAD_UP:                 "VK_GAMEPAD_DPAD_UP",
	VK_GAMEPAD_DPAD_DOWN:               "VK_GAMEPAD_DPAD_DOWN",
	VK_GAMEPAD_DPAD_LEFT:               "VK_GAMEPAD_DPAD_LEFT",
	VK_GAMEPAD_DPAD_RIGHT:              "VK_GAMEPAD_DPAD_RIGHT",
	VK_GAMEPAD_MENU:                    "VK_GAMEPAD_MENU",
	VK_GAMEPAD_VIEW:                    "VK_GAMEPAD_VIEW",
	VK_GAMEPAD_LEFT_THUMBSTICK_BUTTON:  "VK_GAMEPAD_LEFT_THUMBSTICK_BUTTON",
	VK_GAMEPAD_RIGHT_THUMBSTICK_BUTTON: "VK_GAMEPAD_RIGHT_THUMBSTICK_BUTTON",
	VK_GAMEPAD_LEFT_THUMBSTICK_UP:      "VK_GAMEPAD_LEFT_THUMBSTICK_UP",
	VK_GAMEPAD_LEFT_THUMBSTICK_DOWN:    "VK_GAMEPAD_LEFT_THUMBSTICK_DOWN",
	VK_GAMEPAD_LEFT_THUMBSTICK_RIGHT:   "VK_GAMEPAD_LEFT_THUMBSTICK_RIGHT",
	VK_GAMEPAD_LEFT_THUMBSTICK_LEFT:    "VK_GAMEPAD_LEFT_THUMBSTICK_LEFT",
	VK_GAMEPAD_RIGHT_THUMBSTICK_UP:     "VK_GAMEPAD_RIGHT_THUMBSTICK_UP",
	VK_GAMEPAD_RIGHT_THUMBSTICK_DOWN:   "VK_GAMEPAD_RIGHT_THUMBSTICK_DOWN",
	VK_GAMEPAD_RIGHT_THUMBSTICK_RIGHT:  "VK_GAMEPAD_RIGHT_THUMBSTICK_RIGHT",
	VK_GAMEPAD_RIGHT_THUMBSTICK_LEFT:   "VK_GAMEPAD_RIGHT_THUMBSTICK_LEFT",
	VK_OEM_4:                           "VK_OEM_4",
	VK_OEM_5:                           "VK_OEM_5",
	VK_OEM_6:                           "VK_OEM_6",
	VK_OEM_7:                           "VK_OEM_7",
	VK_OEM_8:                           "VK_OEM_8",
	VK_OEM_102:                         "VK_OEM_102",
	VK_PROCESSKEY:                      "VK_PROCESSKEY",
	VK_PACKET:                          "VK_PACKET",
	VK_ATTN:                            "VK_ATTN",
	VK_CRSEL:                           "VK_CRSEL",
	VK_EXSEL:                           "VK_EXSEL",
	VK_EREOF:                           "VK_EREOF",
	VK_PLAY:                            "VK_PLAY",
	VK_ZOOM:                            "VK_ZOOM",
	VK_NONAME:                          "VK_NONAME",
	VK_PA1:                             "VK_PA1",
	VK_OEM_CLEAR:                       "VK_OEM_CLEAR",
}
