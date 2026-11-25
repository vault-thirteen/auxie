package mcci

const (
	Err_UnsupportedAlphaChannelMode = "unsupported alpha channel mode"
	Err_NoChangeIsNeeded            = "no change is needed"
)

type AlphaChannelMode byte

const (
	// AlphaChannelMode_Straight
	// Straight colour values store exact colour values unspoiled by alpha
	// channel. This mode is used for high precision.
	AlphaChannelMode_Straight = AlphaChannelMode(1)

	// AlphaChannelMode_Premultiplied
	// Pre-multiplied colour values store colour values multiplied by alpha
	// value. This mode is used for high speed.
	AlphaChannelMode_Premultiplied = AlphaChannelMode(2)
)

func (acm AlphaChannelMode) IsValid() bool {
	switch acm {
	case AlphaChannelMode_Straight,
		AlphaChannelMode_Premultiplied:
		return true

	default:
		return false
	}
}
