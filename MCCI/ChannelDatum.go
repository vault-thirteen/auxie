package mcci

import "slices"

type ChannelDatum[T IPixel] struct {
	Channel *Channel

	// Canvas is a two-dimensional array of pixels.
	// Rows are counted from bottom to top.
	// Columns are counter from left to right.
	Canvas []Row[T]
}

// CopyCanvas copies the canvas.
func (cd *ChannelDatum[T]) CopyCanvas() (copy []Row[T]) {
	copy = make([]Row[T], len(cd.Canvas))
	for i, row := range cd.Canvas {
		copy[i] = slices.Clone(row)
	}
	return copy
}

// CreateCanvasT creates a new canvas having the same dimensions as the source.
func (cd *ChannelDatum[T]) CreateCanvasT[T2 IPixel]() (newCanvas []Row[T2]) {
	newCanvas = make([]Row[T2], len(cd.Canvas))
	for i, row := range cd.Canvas {
		newCanvas[i] = make([]T2, len(row))
	}
	return newCanvas
}
