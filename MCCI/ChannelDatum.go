package mcci

type ChannelDatum[T IPixel] struct {
	Channel *Channel

	// Canvas is a two-dimensional array of pixels.
	// Rows are counted from bottom to top.
	// Columns are counter from left to right.
	Canvas []Row[T]
}
