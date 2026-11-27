package mcci

import (
	"image/color"
	"path/filepath"
	"testing"

	"github.com/vault-thirteen/auxie/cio"
	"github.com/vault-thirteen/auxie/tester"
)

func Test_ReplaceWithArea(t *testing.T) {
	aTest := tester.New(t)
	filePath := filepath.Join(TestFolder, "3x2.png")

	x, err := NewImageFromFile(filePath, PixelFormat_Integer_1B, AlphaChannelMode_Straight)
	aTest.MustBeNoError(err)
	dstImgI1B, ok := x.(*Image[uint8])
	aTest.MustBeEqual(ok, true)
	cio.PrintMessage(dstImgI1B)

	var srcImgI1B *Image[uint8]
	srcImgI1B, err = dstImgI1B.MakeCopy()
	aTest.MustBeNoError(err)
	srcImgI1B.Resize(2, 1)

	dstArea := NewAreaFromPoints(Point[Size]{X: 1, Y: 1}, Point[Size]{X: 2, Y: 1})
	srcArea := NewAreaFromPoints(Point[Size]{X: 0, Y: 0}, Point[Size]{X: 1, Y: 0})
	err = dstImgI1B.ReplaceWithArea(srcImgI1B, srcArea, dstArea)
	aTest.MustBeNoError(err)
	cio.PrintMessage(dstImgI1B)
}

func Test_ReplaceWithImage(t *testing.T) {
	aTest := tester.New(t)
	filePath := filepath.Join(TestFolder, "3x2.png")

	x, err := NewImageFromFile(filePath, PixelFormat_Integer_1B, AlphaChannelMode_Straight)
	aTest.MustBeNoError(err)
	dstImgI1B, ok := x.(*Image[uint8])
	aTest.MustBeEqual(ok, true)
	cio.PrintMessage(dstImgI1B)

	var srcImgI1B *Image[uint8]
	srcImgI1B, err = dstImgI1B.MakeCopy()
	aTest.MustBeNoError(err)
	srcImgI1B.Resize(2, 1)

	dstArea := NewAreaFromPoints(Point[Size]{X: 1, Y: 1}, Point[Size]{X: 2, Y: 1})
	err = dstImgI1B.ReplaceWithImage(srcImgI1B, dstArea)
	aTest.MustBeNoError(err)
	cio.PrintMessage(dstImgI1B)
}

func Test_ReplaceWithImageAt(t *testing.T) {
	aTest := tester.New(t)
	filePath := filepath.Join(TestFolder, "3x2.png")

	x, err := NewImageFromFile(filePath, PixelFormat_Integer_1B, AlphaChannelMode_Straight)
	aTest.MustBeNoError(err)
	dstImgI1B, ok := x.(*Image[uint8])
	aTest.MustBeEqual(ok, true)
	cio.PrintMessage(dstImgI1B)

	var srcImgI1B *Image[uint8]
	srcImgI1B, err = dstImgI1B.MakeCopy()
	aTest.MustBeNoError(err)
	srcImgI1B.Resize(3, 1)

	p := Point[Size]{X: 0, Y: 1}
	err = dstImgI1B.ReplaceWithImageAt(srcImgI1B, p)
	aTest.MustBeNoError(err)
	cio.PrintMessage(dstImgI1B)
}

func Test_FillWithColour(t *testing.T) {
	aTest := tester.New(t)
	filePath := filepath.Join(TestFolder, "3x2.png")

	x, err := NewImageFromFile(filePath, PixelFormat_Integer_1B, AlphaChannelMode_Straight)
	aTest.MustBeNoError(err)
	dstImgI1B, ok := x.(*Image[uint8])
	aTest.MustBeEqual(ok, true)
	cio.PrintMessage(dstImgI1B)

	colour := color.NRGBA{R: 0x00, G: 0x00, B: 0xFF, A: 0x3F}
	dstArea := NewAreaFromPoints(Point[Size]{X: 0, Y: 0}, Point[Size]{X: 1, Y: 1})
	err = dstImgI1B.FillWithColour(colour, dstArea)
	aTest.MustBeNoError(err)
	cio.PrintMessage(dstImgI1B)
}
