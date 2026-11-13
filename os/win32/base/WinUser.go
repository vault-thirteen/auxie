package b

import (
	"unsafe"
)

// CS constants.
const (
	CS_VREDRAW         = 0x0001
	CS_HREDRAW         = 0x0002
	CS_DBLCLKS         = 0x0008
	CS_OWNDC           = 0x0020
	CS_CLASSDC         = 0x0040
	CS_PARENTDC        = 0x0080
	CS_NOCLOSE         = 0x0200
	CS_SAVEBITS        = 0x0800
	CS_BYTEALIGNCLIENT = 0x1000
	CS_BYTEALIGNWINDOW = 0x2000
	CS_GLOBALCLASS     = 0x4000
	CS_IME             = 0x00010000
	CS_DROPSHADOW      = 0x00020000
)

// SIZE constants.
const (
	SIZE_RESTORED  = 0
	SIZE_MINIMIZED = 1
	SIZE_MAXIMIZED = 2
	SIZE_MAXSHOW   = 3
	SIZE_MAXHIDE   = 4
)

// MakeIntResourceW is a remake of Microsoft's macro MAKEINTRESOURCEW.
// #define MAKEINTRESOURCEW(i) ((LPWSTR)((ULONG_PTR)((WORD)(i))))
func MakeIntResourceW(id uint16) *uint16 {
	return (*uint16)(unsafe.Pointer(uintptr(id)))
}

// IDC constants.
var (
	IDC_ARROW       = MakeIntResourceW(32512)
	IDC_IBEAM       = MakeIntResourceW(32513)
	IDC_WAIT        = MakeIntResourceW(32514)
	IDC_CROSS       = MakeIntResourceW(32515)
	IDC_UPARROW     = MakeIntResourceW(32516)
	IDC_SIZE        = MakeIntResourceW(32640) /* OBSOLETE: use IDC_SIZEALL */
	IDC_ICON        = MakeIntResourceW(32641) /* OBSOLETE: use IDC_ARROW */
	IDC_SIZENWSE    = MakeIntResourceW(32642)
	IDC_SIZENESW    = MakeIntResourceW(32643)
	IDC_SIZEWE      = MakeIntResourceW(32644)
	IDC_SIZENS      = MakeIntResourceW(32645)
	IDC_SIZEALL     = MakeIntResourceW(32646)
	IDC_NO          = MakeIntResourceW(32648) /*not in win3.1 */
	IDC_HAND        = MakeIntResourceW(32649)
	IDC_APPSTARTING = MakeIntResourceW(32650) /*not in win3.1 */
	IDC_HELP        = MakeIntResourceW(32651)
	IDC_PIN         = MakeIntResourceW(32671)
	IDC_PERSON      = MakeIntResourceW(32672)
)

// Color constants.
const (
	COLOR_SCROLLBAR               = 0
	COLOR_BACKGROUND              = 1
	COLOR_ACTIVECAPTION           = 2
	COLOR_INACTIVECAPTION         = 3
	COLOR_MENU                    = 4
	COLOR_WINDOW                  = 5
	COLOR_WINDOWFRAME             = 6
	COLOR_MENUTEXT                = 7
	COLOR_WINDOWTEXT              = 8
	COLOR_CAPTIONTEXT             = 9
	COLOR_ACTIVEBORDER            = 10
	COLOR_INACTIVEBORDER          = 11
	COLOR_APPWORKSPACE            = 12
	COLOR_HIGHLIGHT               = 13
	COLOR_HIGHLIGHTTEXT           = 14
	COLOR_BTNFACE                 = 15
	COLOR_BTNSHADOW               = 16
	COLOR_GRAYTEXT                = 17
	COLOR_BTNTEXT                 = 18
	COLOR_INACTIVECAPTIONTEXT     = 19
	COLOR_BTNHIGHLIGHT            = 20
	COLOR_3DDKSHADOW              = 21
	COLOR_3DLIGHT                 = 22
	COLOR_INFOTEXT                = 23
	COLOR_INFOBK                  = 24
	COLOR_HOTLIGHT                = 26
	COLOR_GRADIENTACTIVECAPTION   = 27
	COLOR_GRADIENTINACTIVECAPTION = 28
	COLOR_MENUHILIGHT             = 29
	COLOR_MENUBAR                 = 30
	COLOR_DESKTOP                 = COLOR_BACKGROUND
	COLOR_3DFACE                  = COLOR_BTNFACE
	COLOR_3DSHADOW                = COLOR_BTNSHADOW
	COLOR_3DHIGHLIGHT             = COLOR_BTNHIGHLIGHT
	COLOR_3DHILIGHT               = COLOR_BTNHIGHLIGHT
	COLOR_BTNHILIGHT              = COLOR_BTNHIGHLIGHT
)

// WS constants.
const (
	WS_OVERLAPPED   = 0x00000000
	WS_POPUP        = 0x80000000
	WS_CHILD        = 0x40000000
	WS_MINIMIZE     = 0x20000000
	WS_VISIBLE      = 0x10000000
	WS_DISABLED     = 0x08000000
	WS_CLIPSIBLINGS = 0x04000000
	WS_CLIPCHILDREN = 0x02000000
	WS_MAXIMIZE     = 0x01000000
	WS_CAPTION      = 0x00C00000
	WS_BORDER       = 0x00800000
	WS_DLGFRAME     = 0x00400000
	WS_VSCROLL      = 0x00200000
	WS_HSCROLL      = 0x00100000
	WS_SYSMENU      = 0x00080000
	WS_THICKFRAME   = 0x00040000
	WS_GROUP        = 0x00020000
	WS_TABSTOP      = 0x00010000
	WS_MINIMIZEBOX  = 0x00020000
	WS_MAXIMIZEBOX  = 0x00010000

	WS_TILED            = WS_OVERLAPPED
	WS_ICONIC           = WS_MINIMIZE
	WS_SIZEBOX          = WS_THICKFRAME
	WS_TILEDWINDOW      = WS_OVERLAPPEDWINDOW
	WS_OVERLAPPEDWINDOW = WS_OVERLAPPED | WS_CAPTION | WS_SYSMENU | WS_THICKFRAME | WS_MINIMIZEBOX | WS_MAXIMIZEBOX
	WS_POPUPWINDOW      = WS_POPUP | WS_BORDER | WS_SYSMENU
	WS_CHILDWINDOW      = WS_CHILD

	WS_EX_DLGMODALFRAME       = 0x00000001
	WS_EX_NOPARENTNOTIFY      = 0x00000004
	WS_EX_TOPMOST             = 0x00000008
	WS_EX_ACCEPTFILES         = 0x00000010
	WS_EX_TRANSPARENT         = 0x00000020
	WS_EX_MDICHILD            = 0x00000040
	WS_EX_TOOLWINDOW          = 0x00000080
	WS_EX_WINDOWEDGE          = 0x00000100
	WS_EX_CLIENTEDGE          = 0x00000200
	WS_EX_CONTEXTHELP         = 0x00000400
	WS_EX_RIGHT               = 0x00001000
	WS_EX_LEFT                = 0x00000000
	WS_EX_RTLREADING          = 0x00002000
	WS_EX_LTRREADING          = 0x00000000
	WS_EX_LEFTSCROLLBAR       = 0x00004000
	WS_EX_RIGHTSCROLLBAR      = 0x00000000
	WS_EX_CONTROLPARENT       = 0x00010000
	WS_EX_STATICEDGE          = 0x00020000
	WS_EX_APPWINDOW           = 0x00040000
	WS_EX_OVERLAPPEDWINDOW    = WS_EX_WINDOWEDGE | WS_EX_CLIENTEDGE
	WS_EX_PALETTEWINDOW       = WS_EX_WINDOWEDGE | WS_EX_TOOLWINDOW | WS_EX_TOPMOST
	WS_EX_LAYERED             = 0x00080000
	WS_EX_NOINHERITLAYOUT     = 0x00100000
	WS_EX_NOREDIRECTIONBITMAP = 0x00200000
	WS_EX_LAYOUTRTL           = 0x00400000
	WS_EX_COMPOSITED          = 0x02000000
	WS_EX_NOACTIVATE          = 0x08000000
)

// CW_USEDEFAULT is a constant for "Create Window".
// #define CW_USEDEFAULT ((int)0x80000000)
const CW_USEDEFAULT int32 = ^0x7FFFFFFF

const (
	RepeatedKeyMarker = 0x40000000 // 1 << 30.
)

// GLWP constants.
const (
	GWLP_WNDPROC    = -4
	GWLP_HINSTANCE  = -6
	GWLP_HWNDPARENT = -8
	GWLP_USERDATA   = -21
	GWLP_ID         = -12
)

// PM constants.
const (
	PM_NOREMOVE = 0x0000
	PM_REMOVE   = 0x0001
	PM_NOYIELD  = 0x0002
)

/* XButton values are WORD flags */
const (
	XBUTTON1 = 0x0001
	XBUTTON2 = 0x0002
)

// VK constants, Virtual Keys.
const (
	VK_LBUTTON                         = 0x01 // Left mouse button
	VK_RBUTTON                         = 0x02 // Right mouse button
	VK_CANCEL                          = 0x03 // Control-break processing
	VK_MBUTTON                         = 0x04 // Middle mouse button
	VK_XBUTTON1                        = 0x05 // X1 mouse button
	VK_XBUTTON2                        = 0x06 // X2 mouse button
	VK_BACK                            = 0x08 // Backspace key
	VK_TAB                             = 0x09 // Tab key
	VK_CLEAR                           = 0x0C // Clear key
	VK_RETURN                          = 0x0D // Enter key
	VK_SHIFT                           = 0x10 // Shift key
	VK_CONTROL                         = 0x11 // Ctrl key
	VK_MENU                            = 0x12 // Alt key
	VK_PAUSE                           = 0x13 // Pause key
	VK_CAPITAL                         = 0x14 // Caps lock key
	VK_KANA                            = 0x15 // IME Kana mode
	VK_HANGUL                          = 0x15 // IME Hangul mode
	VK_IME_ON                          = 0x16 // IME On
	VK_JUNJA                           = 0x17 // IME Junja mode
	VK_FINAL                           = 0x18 // IME final mode
	VK_HANJA                           = 0x19 // IME Hanja mode
	VK_KANJI                           = 0x19 // IME Kanji mode
	VK_IME_OFF                         = 0x1A // IME Off
	VK_ESCAPE                          = 0x1B // Esc key
	VK_CONVERT                         = 0x1C // IME convert
	VK_NONCONVERT                      = 0x1D // IME nonconvert
	VK_ACCEPT                          = 0x1E // IME accept
	VK_MODECHANGE                      = 0x1F // IME mode change request
	VK_SPACE                           = 0x20 // Spacebar key
	VK_PRIOR                           = 0x21 // Page up key
	VK_NEXT                            = 0x22 // Page down key
	VK_END                             = 0x23 // End key
	VK_HOME                            = 0x24 // Home key
	VK_LEFT                            = 0x25 // Left arrow key
	VK_UP                              = 0x26 // Up arrow key
	VK_RIGHT                           = 0x27 // Right arrow key
	VK_DOWN                            = 0x28 // Down arrow key
	VK_SELECT                          = 0x29 // Select key
	VK_PRINT                           = 0x2A // Print key
	VK_EXECUTE                         = 0x2B // Execute key
	VK_SNAPSHOT                        = 0x2C // Print screen key
	VK_INSERT                          = 0x2D // Insert key
	VK_DELETE                          = 0x2E // Delete key
	VK_HELP                            = 0x2F // Help key
	VK_0                               = 0x30 // 0 key
	VK_1                               = 0x31 // 1 key
	VK_2                               = 0x32 // 2 key
	VK_3                               = 0x33 // 3 key
	VK_4                               = 0x34 // 4 key
	VK_5                               = 0x35 // 5 key
	VK_6                               = 0x36 // 6 key
	VK_7                               = 0x37 // 7 key
	VK_8                               = 0x38 // 8 key
	VK_9                               = 0x39 // 9 key
	VK_A                               = 0x41 // A key
	VK_B                               = 0x42 // B key
	VK_C                               = 0x43 // C key
	VK_D                               = 0x44 // D key
	VK_E                               = 0x45 // E key
	VK_F                               = 0x46 // F key
	VK_G                               = 0x47 // G key
	VK_H                               = 0x48 // H key
	VK_I                               = 0x49 // I key
	VK_J                               = 0x4A // J key
	VK_K                               = 0x4B // K key
	VK_L                               = 0x4C // L key
	VK_M                               = 0x4D // M key
	VK_N                               = 0x4E // N key
	VK_O                               = 0x4F // O key
	VK_P                               = 0x50 // P key
	VK_Q                               = 0x51 // Q key
	VK_R                               = 0x52 // R key
	VK_S                               = 0x53 // S key
	VK_T                               = 0x54 // T key
	VK_U                               = 0x55 // U key
	VK_V                               = 0x56 // V key
	VK_W                               = 0x57 // W key
	VK_X                               = 0x58 // X key
	VK_Y                               = 0x59 // Y key
	VK_Z                               = 0x5A // Z key
	VK_LWIN                            = 0x5B // Left Windows logo key
	VK_RWIN                            = 0x5C // Right Windows logo key
	VK_APPS                            = 0x5D // Application key
	VK_SLEEP                           = 0x5F // Computer Sleep key
	VK_NUMPAD0                         = 0x60 // Numeric keypad 0 key
	VK_NUMPAD1                         = 0x61 // Numeric keypad 1 key
	VK_NUMPAD2                         = 0x62 // Numeric keypad 2 key
	VK_NUMPAD3                         = 0x63 // Numeric keypad 3 key
	VK_NUMPAD4                         = 0x64 // Numeric keypad 4 key
	VK_NUMPAD5                         = 0x65 // Numeric keypad 5 key
	VK_NUMPAD6                         = 0x66 // Numeric keypad 6 key
	VK_NUMPAD7                         = 0x67 // Numeric keypad 7 key
	VK_NUMPAD8                         = 0x68 // Numeric keypad 8 key
	VK_NUMPAD9                         = 0x69 // Numeric keypad 9 key
	VK_MULTIPLY                        = 0x6A // Multiply key
	VK_ADD                             = 0x6B // Add key
	VK_SEPARATOR                       = 0x6C // Separator key
	VK_SUBTRACT                        = 0x6D // Subtract key
	VK_DECIMAL                         = 0x6E // Decimal key
	VK_DIVIDE                          = 0x6F // Divide key
	VK_F1                              = 0x70 // F1 key
	VK_F2                              = 0x71 // F2 key
	VK_F3                              = 0x72 // F3 key
	VK_F4                              = 0x73 // F4 key
	VK_F5                              = 0x74 // F5 key
	VK_F6                              = 0x75 // F6 key
	VK_F7                              = 0x76 // F7 key
	VK_F8                              = 0x77 // F8 key
	VK_F9                              = 0x78 // F9 key
	VK_F10                             = 0x79 // F10 key
	VK_F11                             = 0x7A // F11 key
	VK_F12                             = 0x7B // F12 key
	VK_F13                             = 0x7C // F13 key
	VK_F14                             = 0x7D // F14 key
	VK_F15                             = 0x7E // F15 key
	VK_F16                             = 0x7F // F16 key
	VK_F17                             = 0x80 // F17 key
	VK_F18                             = 0x81 // F18 key
	VK_F19                             = 0x82 // F19 key
	VK_F20                             = 0x83 // F20 key
	VK_F21                             = 0x84 // F21 key
	VK_F22                             = 0x85 // F22 key
	VK_F23                             = 0x86 // F23 key
	VK_F24                             = 0x87 // F24 key
	VK_NUMLOCK                         = 0x90 // Num lock key
	VK_SCROLL                          = 0x91 // Scroll lock key
	VK_LSHIFT                          = 0xA0 // Left Shift key
	VK_RSHIFT                          = 0xA1 // Right Shift key
	VK_LCONTROL                        = 0xA2 // Left Ctrl key
	VK_RCONTROL                        = 0xA3 // Right Ctrl key
	VK_LMENU                           = 0xA4 // Left Alt key
	VK_RMENU                           = 0xA5 // Right Alt key
	VK_BROWSER_BACK                    = 0xA6 // Browser Back key
	VK_BROWSER_FORWARD                 = 0xA7 // Browser Forward key
	VK_BROWSER_REFRESH                 = 0xA8 // Browser Refresh key
	VK_BROWSER_STOP                    = 0xA9 // Browser Stop key
	VK_BROWSER_SEARCH                  = 0xAA // Browser Search key
	VK_BROWSER_FAVORITES               = 0xAB // Browser Favorites key
	VK_BROWSER_HOME                    = 0xAC // Browser Start and Home key
	VK_VOLUME_MUTE                     = 0xAD // Volume Mute key
	VK_VOLUME_DOWN                     = 0xAE // Volume Down key
	VK_VOLUME_UP                       = 0xAF // Volume Up key
	VK_MEDIA_NEXT_TRACK                = 0xB0 // Next Track key
	VK_MEDIA_PREV_TRACK                = 0xB1 // Previous Track key
	VK_MEDIA_STOP                      = 0xB2 // Stop Media key
	VK_MEDIA_PLAY_PAUSE                = 0xB3 // Play/Pause Media key
	VK_LAUNCH_MAIL                     = 0xB4 // Start Mail key
	VK_LAUNCH_MEDIA_SELECT             = 0xB5 // Select Media key
	VK_LAUNCH_APP1                     = 0xB6 // Start Application 1 key
	VK_LAUNCH_APP2                     = 0xB7 // Start Application 2 key
	VK_OEM_1                           = 0xBA // It can vary by keyboard. For the US ANSI keyboard, the Semicolon and Colon key
	VK_OEM_PLUS                        = 0xBB // For any country/region, the Equals and Plus key
	VK_OEM_COMMA                       = 0xBC // For any country/region, the Comma and Less Than key
	VK_OEM_MINUS                       = 0xBD // For any country/region, the Dash and Underscore key
	VK_OEM_PERIOD                      = 0xBE // For any country/region, the Period and Greater Than key
	VK_OEM_2                           = 0xBF // It can vary by keyboard. For the US ANSI keyboard, the Forward Slash and Question Mark key
	VK_OEM_3                           = 0xC0 // It can vary by keyboard. For the US ANSI keyboard, the Grave Accent and Tilde key
	VK_GAMEPAD_A                       = 0xC3 // Gamepad A button
	VK_GAMEPAD_B                       = 0xC4 // Gamepad B button
	VK_GAMEPAD_X                       = 0xC5 // Gamepad X button
	VK_GAMEPAD_Y                       = 0xC6 // Gamepad Y button
	VK_GAMEPAD_RIGHT_SHOULDER          = 0xC7 // Gamepad Right Shoulder button
	VK_GAMEPAD_LEFT_SHOULDER           = 0xC8 // Gamepad Left Shoulder button
	VK_GAMEPAD_LEFT_TRIGGER            = 0xC9 // Gamepad Left Trigger button
	VK_GAMEPAD_RIGHT_TRIGGER           = 0xCA // Gamepad Right Trigger button
	VK_GAMEPAD_DPAD_UP                 = 0xCB // Gamepad D-pad Up button
	VK_GAMEPAD_DPAD_DOWN               = 0xCC // Gamepad D-pad Down button
	VK_GAMEPAD_DPAD_LEFT               = 0xCD // Gamepad D-pad Left button
	VK_GAMEPAD_DPAD_RIGHT              = 0xCE // Gamepad D-pad Right button
	VK_GAMEPAD_MENU                    = 0xCF // Gamepad Menu/Start button
	VK_GAMEPAD_VIEW                    = 0xD0 // Gamepad View/Back button
	VK_GAMEPAD_LEFT_THUMBSTICK_BUTTON  = 0xD1 // Gamepad Left Thumbstick button
	VK_GAMEPAD_RIGHT_THUMBSTICK_BUTTON = 0xD2 // Gamepad Right Thumbstick button
	VK_GAMEPAD_LEFT_THUMBSTICK_UP      = 0xD3 // Gamepad Left Thumbstick up
	VK_GAMEPAD_LEFT_THUMBSTICK_DOWN    = 0xD4 // Gamepad Left Thumbstick down
	VK_GAMEPAD_LEFT_THUMBSTICK_RIGHT   = 0xD5 // Gamepad Left Thumbstick right
	VK_GAMEPAD_LEFT_THUMBSTICK_LEFT    = 0xD6 // Gamepad Left Thumbstick left
	VK_GAMEPAD_RIGHT_THUMBSTICK_UP     = 0xD7 // Gamepad Right Thumbstick up
	VK_GAMEPAD_RIGHT_THUMBSTICK_DOWN   = 0xD8 // Gamepad Right Thumbstick down
	VK_GAMEPAD_RIGHT_THUMBSTICK_RIGHT  = 0xD9 // Gamepad Right Thumbstick right
	VK_GAMEPAD_RIGHT_THUMBSTICK_LEFT   = 0xDA // Gamepad Right Thumbstick left
	VK_OEM_4                           = 0xDB // It can vary by keyboard. For the US ANSI keyboard, the Left Brace key
	VK_OEM_5                           = 0xDC // It can vary by keyboard. For the US ANSI keyboard, the Backslash and Pipe key
	VK_OEM_6                           = 0xDD // It can vary by keyboard. For the US ANSI keyboard, the Right Brace key
	VK_OEM_7                           = 0xDE // It can vary by keyboard. For the US ANSI keyboard, the Apostrophe and Double Quotation Mark key
	VK_OEM_8                           = 0xDF // It can vary by keyboard. For the Canadian CSA keyboard, the Right Ctrl key
	VK_OEM_102                         = 0xE2 // It can vary by keyboard. For the European ISO keyboard, the Backslash and Pipe key
	VK_PROCESSKEY                      = 0xE5 // IME PROCESS key
	VK_PACKET                          = 0xE7 // Used to pass Unicode characters as if they were keystrokes. The VK_PACKET key is the low word of a 32-bit Virtual Key value used for non-keyboard input methods. For more information, see Remark in KEYBDINPUT, SendInput, WM_KEYDOWN, and WM_KEYUP
	VK_ATTN                            = 0xF6 // Attn key
	VK_CRSEL                           = 0xF7 // CrSel key
	VK_EXSEL                           = 0xF8 // ExSel key
	VK_EREOF                           = 0xF9 // Erase EOF key
	VK_PLAY                            = 0xFA // Play key
	VK_ZOOM                            = 0xFB // Zoom key
	VK_NONAME                          = 0xFC // Reserved
	VK_PA1                             = 0xFD // PA1 key
	VK_OEM_CLEAR                       = 0xFE // Clear key
)
