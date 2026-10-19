package cio

// KeyCode is a virtual key code.
// This functionality is guaranteed to work only on Windows operating systems.
type KeyCode byte

// Virtual Key codes.
// Source:
// https://learn.microsoft.com/en-gb/windows/win32/inputdev/virtual-key-codes
const (
	VK_LBUTTON                         KeyCode = 0x01 // Left mouse button
	VK_RBUTTON                         KeyCode = 0x02 // Right mouse button
	VK_CANCEL                          KeyCode = 0x03 // Control-break processing
	VK_MBUTTON                         KeyCode = 0x04 // Middle mouse button
	VK_XBUTTON1                        KeyCode = 0x05 // X1 mouse button
	VK_XBUTTON2                        KeyCode = 0x06 // X2 mouse button
	VK_BACK                            KeyCode = 0x08 // Backspace key
	VK_TAB                             KeyCode = 0x09 // Tab key
	VK_CLEAR                           KeyCode = 0x0C // Clear key
	VK_RETURN                          KeyCode = 0x0D // Enter key
	VK_SHIFT                           KeyCode = 0x10 // Shift key
	VK_CONTROL                         KeyCode = 0x11 // Ctrl key
	VK_MENU                            KeyCode = 0x12 // Alt key
	VK_PAUSE                           KeyCode = 0x13 // Pause key
	VK_CAPITAL                         KeyCode = 0x14 // Caps lock key
	VK_KANA                            KeyCode = 0x15 // IME Kana mode
	VK_HANGUL                          KeyCode = 0x15 // IME Hangul mode
	VK_IME_ON                          KeyCode = 0x16 // IME On
	VK_JUNJA                           KeyCode = 0x17 // IME Junja mode
	VK_FINAL                           KeyCode = 0x18 // IME final mode
	VK_HANJA                           KeyCode = 0x19 // IME Hanja mode
	VK_KANJI                           KeyCode = 0x19 // IME Kanji mode
	VK_IME_OFF                         KeyCode = 0x1A // IME Off
	VK_ESCAPE                          KeyCode = 0x1B // Esc key
	VK_CONVERT                         KeyCode = 0x1C // IME convert
	VK_NONCONVERT                      KeyCode = 0x1D // IME nonconvert
	VK_ACCEPT                          KeyCode = 0x1E // IME accept
	VK_MODECHANGE                      KeyCode = 0x1F // IME mode change request
	VK_SPACE                           KeyCode = 0x20 // Spacebar key
	VK_PRIOR                           KeyCode = 0x21 // Page up key
	VK_NEXT                            KeyCode = 0x22 // Page down key
	VK_END                             KeyCode = 0x23 // End key
	VK_HOME                            KeyCode = 0x24 // Home key
	VK_LEFT                            KeyCode = 0x25 // Left arrow key
	VK_UP                              KeyCode = 0x26 // Up arrow key
	VK_RIGHT                           KeyCode = 0x27 // Right arrow key
	VK_DOWN                            KeyCode = 0x28 // Down arrow key
	VK_SELECT                          KeyCode = 0x29 // Select key
	VK_PRINT                           KeyCode = 0x2A // Print key
	VK_EXECUTE                         KeyCode = 0x2B // Execute key
	VK_SNAPSHOT                        KeyCode = 0x2C // Print screen key
	VK_INSERT                          KeyCode = 0x2D // Insert key
	VK_DELETE                          KeyCode = 0x2E // Delete key
	VK_HELP                            KeyCode = 0x2F // Help key
	VK_0                               KeyCode = 0x30 // 0 key
	VK_1                               KeyCode = 0x31 // 1 key
	VK_2                               KeyCode = 0x32 // 2 key
	VK_3                               KeyCode = 0x33 // 3 key
	VK_4                               KeyCode = 0x34 // 4 key
	VK_5                               KeyCode = 0x35 // 5 key
	VK_6                               KeyCode = 0x36 // 6 key
	VK_7                               KeyCode = 0x37 // 7 key
	VK_8                               KeyCode = 0x38 // 8 key
	VK_9                               KeyCode = 0x39 // 9 key
	VK_A                               KeyCode = 0x41 // A key
	VK_B                               KeyCode = 0x42 // B key
	VK_C                               KeyCode = 0x43 // C key
	VK_D                               KeyCode = 0x44 // D key
	VK_E                               KeyCode = 0x45 // E key
	VK_F                               KeyCode = 0x46 // F key
	VK_G                               KeyCode = 0x47 // G key
	VK_H                               KeyCode = 0x48 // H key
	VK_I                               KeyCode = 0x49 // I key
	VK_J                               KeyCode = 0x4A // J key
	VK_K                               KeyCode = 0x4B // K key
	VK_L                               KeyCode = 0x4C // L key
	VK_M                               KeyCode = 0x4D // M key
	VK_N                               KeyCode = 0x4E // N key
	VK_O                               KeyCode = 0x4F // O key
	VK_P                               KeyCode = 0x50 // P key
	VK_Q                               KeyCode = 0x51 // Q key
	VK_R                               KeyCode = 0x52 // R key
	VK_S                               KeyCode = 0x53 // S key
	VK_T                               KeyCode = 0x54 // T key
	VK_U                               KeyCode = 0x55 // U key
	VK_V                               KeyCode = 0x56 // V key
	VK_W                               KeyCode = 0x57 // W key
	VK_X                               KeyCode = 0x58 // X key
	VK_Y                               KeyCode = 0x59 // Y key
	VK_Z                               KeyCode = 0x5A // Z key
	VK_LWIN                            KeyCode = 0x5B // Left Windows logo key
	VK_RWIN                            KeyCode = 0x5C // Right Windows logo key
	VK_APPS                            KeyCode = 0x5D // Application key
	VK_SLEEP                           KeyCode = 0x5F // Computer Sleep key
	VK_NUMPAD0                         KeyCode = 0x60 // Numeric keypad 0 key
	VK_NUMPAD1                         KeyCode = 0x61 // Numeric keypad 1 key
	VK_NUMPAD2                         KeyCode = 0x62 // Numeric keypad 2 key
	VK_NUMPAD3                         KeyCode = 0x63 // Numeric keypad 3 key
	VK_NUMPAD4                         KeyCode = 0x64 // Numeric keypad 4 key
	VK_NUMPAD5                         KeyCode = 0x65 // Numeric keypad 5 key
	VK_NUMPAD6                         KeyCode = 0x66 // Numeric keypad 6 key
	VK_NUMPAD7                         KeyCode = 0x67 // Numeric keypad 7 key
	VK_NUMPAD8                         KeyCode = 0x68 // Numeric keypad 8 key
	VK_NUMPAD9                         KeyCode = 0x69 // Numeric keypad 9 key
	VK_MULTIPLY                        KeyCode = 0x6A // Multiply key
	VK_ADD                             KeyCode = 0x6B // Add key
	VK_SEPARATOR                       KeyCode = 0x6C // Separator key
	VK_SUBTRACT                        KeyCode = 0x6D // Subtract key
	VK_DECIMAL                         KeyCode = 0x6E // Decimal key
	VK_DIVIDE                          KeyCode = 0x6F // Divide key
	VK_F1                              KeyCode = 0x70 // F1 key
	VK_F2                              KeyCode = 0x71 // F2 key
	VK_F3                              KeyCode = 0x72 // F3 key
	VK_F4                              KeyCode = 0x73 // F4 key
	VK_F5                              KeyCode = 0x74 // F5 key
	VK_F6                              KeyCode = 0x75 // F6 key
	VK_F7                              KeyCode = 0x76 // F7 key
	VK_F8                              KeyCode = 0x77 // F8 key
	VK_F9                              KeyCode = 0x78 // F9 key
	VK_F10                             KeyCode = 0x79 // F10 key
	VK_F11                             KeyCode = 0x7A // F11 key
	VK_F12                             KeyCode = 0x7B // F12 key
	VK_F13                             KeyCode = 0x7C // F13 key
	VK_F14                             KeyCode = 0x7D // F14 key
	VK_F15                             KeyCode = 0x7E // F15 key
	VK_F16                             KeyCode = 0x7F // F16 key
	VK_F17                             KeyCode = 0x80 // F17 key
	VK_F18                             KeyCode = 0x81 // F18 key
	VK_F19                             KeyCode = 0x82 // F19 key
	VK_F20                             KeyCode = 0x83 // F20 key
	VK_F21                             KeyCode = 0x84 // F21 key
	VK_F22                             KeyCode = 0x85 // F22 key
	VK_F23                             KeyCode = 0x86 // F23 key
	VK_F24                             KeyCode = 0x87 // F24 key
	VK_NUMLOCK                         KeyCode = 0x90 // Num lock key
	VK_SCROLL                          KeyCode = 0x91 // Scroll lock key
	VK_LSHIFT                          KeyCode = 0xA0 // Left Shift key
	VK_RSHIFT                          KeyCode = 0xA1 // Right Shift key
	VK_LCONTROL                        KeyCode = 0xA2 // Left Ctrl key
	VK_RCONTROL                        KeyCode = 0xA3 // Right Ctrl key
	VK_LMENU                           KeyCode = 0xA4 // Left Alt key
	VK_RMENU                           KeyCode = 0xA5 // Right Alt key
	VK_BROWSER_BACK                    KeyCode = 0xA6 // Browser Back key
	VK_BROWSER_FORWARD                 KeyCode = 0xA7 // Browser Forward key
	VK_BROWSER_REFRESH                 KeyCode = 0xA8 // Browser Refresh key
	VK_BROWSER_STOP                    KeyCode = 0xA9 // Browser Stop key
	VK_BROWSER_SEARCH                  KeyCode = 0xAA // Browser Search key
	VK_BROWSER_FAVORITES               KeyCode = 0xAB // Browser Favorites key
	VK_BROWSER_HOME                    KeyCode = 0xAC // Browser Start and Home key
	VK_VOLUME_MUTE                     KeyCode = 0xAD // Volume Mute key
	VK_VOLUME_DOWN                     KeyCode = 0xAE // Volume Down key
	VK_VOLUME_UP                       KeyCode = 0xAF // Volume Up key
	VK_MEDIA_NEXT_TRACK                KeyCode = 0xB0 // Next Track key
	VK_MEDIA_PREV_TRACK                KeyCode = 0xB1 // Previous Track key
	VK_MEDIA_STOP                      KeyCode = 0xB2 // Stop Media key
	VK_MEDIA_PLAY_PAUSE                KeyCode = 0xB3 // Play/Pause Media key
	VK_LAUNCH_MAIL                     KeyCode = 0xB4 // Start Mail key
	VK_LAUNCH_MEDIA_SELECT             KeyCode = 0xB5 // Select Media key
	VK_LAUNCH_APP1                     KeyCode = 0xB6 // Start Application 1 key
	VK_LAUNCH_APP2                     KeyCode = 0xB7 // Start Application 2 key
	VK_OEM_1                           KeyCode = 0xBA // It can vary by keyboard. For the US ANSI keyboard, the Semicolon and Colon key
	VK_OEM_PLUS                        KeyCode = 0xBB // For any country/region, the Equals and Plus key
	VK_OEM_COMMA                       KeyCode = 0xBC // For any country/region, the Comma and Less Than key
	VK_OEM_MINUS                       KeyCode = 0xBD // For any country/region, the Dash and Underscore key
	VK_OEM_PERIOD                      KeyCode = 0xBE // For any country/region, the Period and Greater Than key
	VK_OEM_2                           KeyCode = 0xBF // It can vary by keyboard. For the US ANSI keyboard, the Forward Slash and Question Mark key
	VK_OEM_3                           KeyCode = 0xC0 // It can vary by keyboard. For the US ANSI keyboard, the Grave Accent and Tilde key
	VK_GAMEPAD_A                       KeyCode = 0xC3 // Gamepad A button
	VK_GAMEPAD_B                       KeyCode = 0xC4 // Gamepad B button
	VK_GAMEPAD_X                       KeyCode = 0xC5 // Gamepad X button
	VK_GAMEPAD_Y                       KeyCode = 0xC6 // Gamepad Y button
	VK_GAMEPAD_RIGHT_SHOULDER          KeyCode = 0xC7 // Gamepad Right Shoulder button
	VK_GAMEPAD_LEFT_SHOULDER           KeyCode = 0xC8 // Gamepad Left Shoulder button
	VK_GAMEPAD_LEFT_TRIGGER            KeyCode = 0xC9 // Gamepad Left Trigger button
	VK_GAMEPAD_RIGHT_TRIGGER           KeyCode = 0xCA // Gamepad Right Trigger button
	VK_GAMEPAD_DPAD_UP                 KeyCode = 0xCB // Gamepad D-pad Up button
	VK_GAMEPAD_DPAD_DOWN               KeyCode = 0xCC // Gamepad D-pad Down button
	VK_GAMEPAD_DPAD_LEFT               KeyCode = 0xCD // Gamepad D-pad Left button
	VK_GAMEPAD_DPAD_RIGHT              KeyCode = 0xCE // Gamepad D-pad Right button
	VK_GAMEPAD_MENU                    KeyCode = 0xCF // Gamepad Menu/Start button
	VK_GAMEPAD_VIEW                    KeyCode = 0xD0 // Gamepad View/Back button
	VK_GAMEPAD_LEFT_THUMBSTICK_BUTTON  KeyCode = 0xD1 // Gamepad Left Thumbstick button
	VK_GAMEPAD_RIGHT_THUMBSTICK_BUTTON KeyCode = 0xD2 // Gamepad Right Thumbstick button
	VK_GAMEPAD_LEFT_THUMBSTICK_UP      KeyCode = 0xD3 // Gamepad Left Thumbstick up
	VK_GAMEPAD_LEFT_THUMBSTICK_DOWN    KeyCode = 0xD4 // Gamepad Left Thumbstick down
	VK_GAMEPAD_LEFT_THUMBSTICK_RIGHT   KeyCode = 0xD5 // Gamepad Left Thumbstick right
	VK_GAMEPAD_LEFT_THUMBSTICK_LEFT    KeyCode = 0xD6 // Gamepad Left Thumbstick left
	VK_GAMEPAD_RIGHT_THUMBSTICK_UP     KeyCode = 0xD7 // Gamepad Right Thumbstick up
	VK_GAMEPAD_RIGHT_THUMBSTICK_DOWN   KeyCode = 0xD8 // Gamepad Right Thumbstick down
	VK_GAMEPAD_RIGHT_THUMBSTICK_RIGHT  KeyCode = 0xD9 // Gamepad Right Thumbstick right
	VK_GAMEPAD_RIGHT_THUMBSTICK_LEFT   KeyCode = 0xDA // Gamepad Right Thumbstick left
	VK_OEM_4                           KeyCode = 0xDB // It can vary by keyboard. For the US ANSI keyboard, the Left Brace key
	VK_OEM_5                           KeyCode = 0xDC // It can vary by keyboard. For the US ANSI keyboard, the Backslash and Pipe key
	VK_OEM_6                           KeyCode = 0xDD // It can vary by keyboard. For the US ANSI keyboard, the Right Brace key
	VK_OEM_7                           KeyCode = 0xDE // It can vary by keyboard. For the US ANSI keyboard, the Apostrophe and Double Quotation Mark key
	VK_OEM_8                           KeyCode = 0xDF // It can vary by keyboard. For the Canadian CSA keyboard, the Right Ctrl key
	VK_OEM_102                         KeyCode = 0xE2 // It can vary by keyboard. For the European ISO keyboard, the Backslash and Pipe key
	VK_PROCESSKEY                      KeyCode = 0xE5 // IME PROCESS key
	VK_PACKET                          KeyCode = 0xE7 // Used to pass Unicode characters as if they were keystrokes. The VK_PACKET key is the low word of a 32-bit Virtual Key value used for non-keyboard input methods. For more information, see Remark in KEYBDINPUT, SendInput, WM_KEYDOWN, and WM_KEYUP
	VK_ATTN                            KeyCode = 0xF6 // Attn key
	VK_CRSEL                           KeyCode = 0xF7 // CrSel key
	VK_EXSEL                           KeyCode = 0xF8 // ExSel key
	VK_EREOF                           KeyCode = 0xF9 // Erase EOF key
	VK_PLAY                            KeyCode = 0xFA // Play key
	VK_ZOOM                            KeyCode = 0xFB // Zoom key
	VK_NONAME                          KeyCode = 0xFC // Reserved
	VK_PA1                             KeyCode = 0xFD // PA1 key
	VK_OEM_CLEAR                       KeyCode = 0xFE // Clear key
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
