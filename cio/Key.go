package cio

import (
	"bytes"

	"github.com/vault-thirteen/auxie/errors"
)

const (
	Err_InvalidKey = "invalid key"
	Err_UnknownKey = "unknown key"
)

// Key is a keyboard key pressed in a terminal console.
// This functionality is guaranteed to work only on Windows operating systems.
type Key struct {
	bytes  []byte
	vkCode KeyCode
	name   string
}

var singleByteKeys = map[byte]KeyCode{
	// Functional keys row.
	27: VK_ESCAPE, // Escape.

	// First row.
	96:  VK_OEM_3,     // Tilde.
	49:  VK_1,         // 1.
	50:  VK_2,         // 2.
	51:  VK_3,         // 3.
	52:  VK_4,         // 4.
	53:  VK_5,         // 5.
	54:  VK_6,         // 6.
	55:  VK_7,         //7.
	56:  VK_8,         //8.
	57:  VK_9,         //9.
	48:  VK_0,         // 0.
	45:  VK_OEM_MINUS, // 3.
	61:  VK_OEM_PLUS,  // =.
	127: VK_BACK,      // Backspace.

	// First row with Shift.
	126: VK_OEM_3,     // Backtick.
	33:  VK_1,         // !.
	64:  VK_2,         // @.
	35:  VK_3,         // #.
	36:  VK_4,         // $.
	37:  VK_5,         // %.
	94:  VK_6,         // ^.
	38:  VK_7,         //&.
	42:  VK_8,         //*.
	40:  VK_9,         //(.
	41:  VK_0,         // ).
	95:  VK_OEM_MINUS, // _.
	43:  VK_OEM_PLUS,  // +.

	// Second row.
	9:   VK_TAB,   // Tab.
	113: VK_Q,     // q.
	119: VK_W,     // w.
	101: VK_E,     // e.
	114: VK_R,     // r.
	116: VK_T,     // t.
	121: VK_Y,     // y.
	117: VK_U,     // u.
	105: VK_I,     // i.
	111: VK_O,     // o.
	112: VK_P,     // p.
	91:  VK_OEM_4, // [.
	93:  VK_OEM_6, // ].
	92:  VK_OEM_5, // \.

	// Second row with Shift
	81:  VK_Q,     // Q.
	87:  VK_W,     // W.
	69:  VK_E,     // E.
	82:  VK_R,     // R.
	84:  VK_T,     // T.
	89:  VK_Y,     // Y.
	85:  VK_U,     // U.
	73:  VK_I,     // I.
	79:  VK_O,     // O.
	80:  VK_P,     // P.
	123: VK_OEM_4, // {.
	125: VK_OEM_6, // }.
	124: VK_OEM_5, // |.

	// Third row.
	97:  VK_A,      // a.
	115: VK_S,      // s.
	100: VK_D,      // d.
	102: VK_F,      // f.
	103: VK_G,      // g.
	104: VK_H,      // h.
	106: VK_J,      // j.
	107: VK_K,      // k.
	108: VK_L,      // l.
	59:  VK_OEM_1,  // ;.
	39:  VK_OEM_7,  // '.
	13:  VK_RETURN, // Enter.

	// Third row with Shift.
	65: VK_A,     // A.
	83: VK_S,     // S.
	68: VK_D,     // D.
	70: VK_F,     // F.
	71: VK_G,     // G.
	72: VK_H,     // H.
	74: VK_J,     // J.
	75: VK_K,     // K.
	76: VK_L,     // L.
	58: VK_OEM_1, // :.
	34: VK_OEM_7, // ".

	// Fourth row.
	122: VK_Z,          // z.
	120: VK_X,          // x.
	99:  VK_C,          // c.
	118: VK_V,          // v.
	98:  VK_B,          // b.
	110: VK_N,          // n.
	109: VK_M,          // m.
	44:  VK_OEM_COMMA,  // ,.
	46:  VK_OEM_PERIOD, // ..
	47:  VK_OEM_2,      // /.

	// Fourth row with Shift.
	90: VK_Z,          // Z.
	88: VK_X,          // X.
	67: VK_C,          // C.
	86: VK_V,          // V.
	66: VK_B,          // B.
	78: VK_N,          // N.
	77: VK_M,          // M.
	60: VK_OEM_COMMA,  // <.
	62: VK_OEM_PERIOD, // >.
	63: VK_OEM_2,      // ?.

	// Bottom row.
	32: VK_SPACE, // Spacebar.
}

var (
	byteSeq_F1 = []byte{27, 79, 80}
	byteSeq_F2 = []byte{27, 79, 81}
	byteSeq_F3 = []byte{27, 79, 82}
	byteSeq_F4 = []byte{27, 79, 83}

	byteSeq_F5 = []byte{27, 91, 49, 53, 126}
	byteSeq_F6 = []byte{27, 91, 49, 55, 126}
	byteSeq_F7 = []byte{27, 91, 49, 56, 126}
	byteSeq_F8 = []byte{27, 91, 49, 57, 126}

	byteSeq_F9  = []byte{27, 91, 50, 48, 126}
	byteSeq_F10 = []byte{27, 91, 50, 49, 126}
	byteSeq_F11 = []byte{27, 91, 50, 51, 126}
	byteSeq_F12 = []byte{27, 91, 50, 52, 126}

	byteSeq_Insert = []byte{27, 91, 50, 126}
	byteSeq_Delete = []byte{27, 91, 51, 126}

	byteSeq_Home = []byte{27, 91, 72}
	byteSeq_End  = []byte{27, 91, 70}

	byteSeq_PageUp   = []byte{27, 91, 53, 126}
	byteSeq_PageDown = []byte{27, 91, 54, 126}

	byteSeq_ArrowLeft  = []byte{27, 91, 68}
	byteSeq_ArrowRight = []byte{27, 91, 67}
	byteSeq_ArrowUp    = []byte{27, 91, 65}
	byteSeq_ArrowDown  = []byte{27, 91, 66}
)

func NewKey(seq []byte, code KeyCode) Key {
	return Key{seq, code, ""}
}

func (k Key) Code() KeyCode {
	return k.vkCode
}

func (k Key) Name() string {
	if len(k.name) == 0 {
		k.name = KeyCodeNames[k.vkCode]
	}

	return k.name
}

func KeyFromTerminalByteSequence(seq []byte) (key Key, err error) {
	switch len(seq) {
	case 0:
		return key, errors.ErrorFromText(Err_InvalidKey)

	case 1:
		{
			code, ok := singleByteKeys[seq[0]]
			if !ok {
				return key, errors.ErrorFromText(Err_UnknownKey)
			}

			return NewKey(seq, code), nil
		}

	case 3:
		{
			if bytes.Equal(seq, byteSeq_F1) {
				return NewKey(seq, VK_F1), nil
			}
			if bytes.Equal(seq, byteSeq_F2) {
				return NewKey(seq, VK_F2), nil
			}
			if bytes.Equal(seq, byteSeq_F3) {
				return NewKey(seq, VK_F3), nil
			}
			if bytes.Equal(seq, byteSeq_F4) {
				return NewKey(seq, VK_F4), nil
			}
			if bytes.Equal(seq, byteSeq_Home) {
				return NewKey(seq, VK_HOME), nil
			}
			if bytes.Equal(seq, byteSeq_End) {
				return NewKey(seq, VK_END), nil
			}
			if bytes.Equal(seq, byteSeq_ArrowLeft) {
				return NewKey(seq, VK_LEFT), nil
			}
			if bytes.Equal(seq, byteSeq_ArrowRight) {
				return NewKey(seq, VK_RIGHT), nil
			}
			if bytes.Equal(seq, byteSeq_ArrowUp) {
				return NewKey(seq, VK_UP), nil
			}
			if bytes.Equal(seq, byteSeq_ArrowDown) {
				return NewKey(seq, VK_DOWN), nil
			}
		}

	case 4:
		{
			if bytes.Equal(seq, byteSeq_Insert) {
				return NewKey(seq, VK_INSERT), nil
			}
			if bytes.Equal(seq, byteSeq_Delete) {
				return NewKey(seq, VK_DELETE), nil
			}
			if bytes.Equal(seq, byteSeq_PageUp) {
				return NewKey(seq, VK_PRIOR), nil
			}
			if bytes.Equal(seq, byteSeq_PageDown) {
				return NewKey(seq, VK_NEXT), nil
			}
		}

	case 5:
		{
			if bytes.Equal(seq, byteSeq_F5) {
				return NewKey(seq, VK_F5), nil
			}
			if bytes.Equal(seq, byteSeq_F6) {
				return NewKey(seq, VK_F6), nil
			}
			if bytes.Equal(seq, byteSeq_F7) {
				return NewKey(seq, VK_F7), nil
			}
			if bytes.Equal(seq, byteSeq_F8) {
				return NewKey(seq, VK_F8), nil
			}
			if bytes.Equal(seq, byteSeq_F9) {
				return NewKey(seq, VK_F9), nil
			}
			if bytes.Equal(seq, byteSeq_F10) {
				return NewKey(seq, VK_F10), nil
			}
			if bytes.Equal(seq, byteSeq_F11) {
				return NewKey(seq, VK_F11), nil
			}
			if bytes.Equal(seq, byteSeq_F12) {
				return NewKey(seq, VK_F12), nil
			}
		}

	default:
		return key, errors.ErrorFromText(Err_InvalidKey)
	}

	return key, nil
}
