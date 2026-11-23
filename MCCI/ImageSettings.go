package mcci

type ImageSettings struct {
	// Revision (version) of file format.
	FileFormatRevision string

	// Width in pixels.
	Width uint16

	// Height in pixels.
	Height uint16

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
