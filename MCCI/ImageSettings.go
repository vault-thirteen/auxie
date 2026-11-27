package mcci

type ImageSettings struct {
	// Revision (version) of file format.
	FileFormatRevision string

	// Width in pixels.
	Width Size

	// Height in pixels.
	Height Size

	// Number of colour channels, including alpha channel if present.
	ChannelCount byte

	// Format of pixels of each colour channel.
	ChannelPixelFormat ChannelPixelFormat

	// Alpha channel mode.
	AlphaChannelMode AlphaChannelMode
}

const (
	FileFormatRevision_1 = "1"
)
