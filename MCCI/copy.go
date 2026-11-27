package mcci

import (
	"slices"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// MakeCopy copies the object.
func (img *Image[T]) MakeCopy() (copy *Image[T], err error) {
	// Step I. Copy simple fields.
	copy = &Image[T]{
		Settings:     img.copySettings(),
		Channels:     img.copyChannels(),
		ChannelIndex: img.copyChannelIndex(),
	}

	// Step II. Copy the 'ChannelDatum' field recreating pointers.
	err = img.copyChannelData(copy)
	if err != nil {
		return nil, err
	}

	return copy, nil
}
func (img *Image[T]) copySettings() (copy *ImageSettings) {
	obj := *img.Settings
	return &obj
}
func (img *Image[T]) copyChannels() (copy []Channel) {
	return slices.Clone(img.Channels)
}
func (img *Image[T]) copyChannelIndex() (copy ChannelIndex) {
	return img.ChannelIndex
}

// copyChannelData copies the 'ChannelData' field into the target object 'obj'.
// For this method to work, all other fields must be already present in the
// target object.
func (img *Image[T]) copyChannelData(obj *Image[T]) (err error) {
	var cópy = make([]ChannelDatum[T], len(img.ChannelData))
	var idx int
	for i, cd := range img.ChannelData {
		idx, err = obj.findChannelIndex(cd.Channel.Name, cd.Channel.WaveLength_Nm)
		if err != nil {
			return err
		}

		cópy[i] = ChannelDatum[T]{
			Channel: &obj.Channels[idx], // Recreate the pointer by index.
			Canvas:  cd.CopyCanvas(),
		}
	}
	obj.ChannelData = cópy
	return nil
}

// MakeCopyT copies the object using a new generic type, except the pixels.
// Canvas pixels are not copied here as this process requires calculations.
func (img *Image[T]) MakeCopyT[T2 IPixel]() (copy *Image[T2], err error) {
	// Step I. Copy simple fields.
	copy = &Image[T2]{
		Settings:     img.copySettings(),
		Channels:     img.copyChannels(),
		ChannelIndex: img.copyChannelIndex(),
	}

	// Step II. Recreate the 'ChannelDatum' field with pointers and empty canvas.
	err = img.createChannelDataT[T2](copy)
	if err != nil {
		return nil, err
	}

	return copy, nil
}

// createChannelDataT recreates the 'ChannelData' field into the target object
// 'obj'. For this method to work, all other fields must be already present in
// the target object. Canvas is not fully copied, only dimensions are copied.
func (img *Image[T]) createChannelDataT[T2 IPixel](obj *Image[T2]) (err error) {
	var cópy = make([]ChannelDatum[T2], len(img.ChannelData))
	var idx int
	for i, cd := range img.ChannelData {
		idx, err = obj.findChannelIndex(cd.Channel.Name, cd.Channel.WaveLength_Nm)
		if err != nil {
			return err
		}

		cópy[i] = ChannelDatum[T2]{
			Channel: &obj.Channels[idx],     // Recreate the pointer by index.
			Canvas:  cd.CreateCanvasT[T2](), // Empty canvas with new type.
		}
	}
	obj.ChannelData = cópy
	return nil
}

// Resize changes canvas size of the image.
func (img *Image[T]) Resize(width, height uint16) {
	for _, cd := range img.ChannelData {
		cd.ResizeCanvas(width, height)
	}

	img.Settings.Width = width
	img.Settings.Height = height
}

// CloneArrayOfNumbers clones an array of numbers into an array of a new type.
// New type must be compatible with the old one. Do not use this function if
// the target type is narrower than original type.
func CloneArrayOfNumbers[T1 Number, T2 Number](in []T1) (out []T2) {
	if in == nil {
		return nil
	}

	out = make([]T2, len(in))
	for i, v := range in {
		out[i] = T2(v)
	}

	return out
}
