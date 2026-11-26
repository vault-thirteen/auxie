package mcci

import (
	"slices"
)

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

// ResizeCanvas resizes the canvas.
// Old pixels within the new dimensions are not destroyed.
// This method, as opposed to standard Go language resizing, frees the memory.
func (cd *ChannelDatum[T]) ResizeCanvas(width, height uint16) {
	if (width <= 0) || (height <= 0) {
		cd.Canvas = nil
		return
	}

	newCanvas := make([]Row[T], height)
	for i := 0; i < int(height); i++ {
		newCanvas[i] = make([]T, width)
	}

	// Do not copy anything if canvas had zero dimensions.
	if len(cd.Canvas) <= 0 {
		cd.Canvas = newCanvas
		return
	}
	if len(cd.Canvas[0]) <= 0 {
		cd.Canvas = newCanvas
		return
	}

	// Copy old pixels within new dimensions.
	var yMax = int(min(height, uint16(len(cd.Canvas))) - 1)
	var xMax = int(min(width, uint16(len(cd.Canvas[0]))) - 1)

	for y := 0; y <= yMax; y++ {
		for x := 0; x <= xMax; x++ {
			newCanvas[y][x] = cd.Canvas[y][x]
		}
	}

	// While garbage collector in Go language is tracing and not
	// reference-counted, old slice data will be deleted some time later.
	cd.Canvas = newCanvas
	return
}
