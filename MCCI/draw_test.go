package mcci

import (
	"path/filepath"
	"testing"

	"github.com/vault-thirteen/auxie/cio"
	"github.com/vault-thirteen/auxie/tester"
)

func Test_DrawOverArea(t *testing.T) {
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
	err = dstImgI1B.DrawOverArea(srcImgI1B, srcArea, dstArea)
	aTest.MustBeNoError(err)
	cio.PrintMessage(dstImgI1B)
}

func Test_DrawOverImage(t *testing.T) {
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
	err = dstImgI1B.DrawOverImage(srcImgI1B, dstArea)
	aTest.MustBeNoError(err)
	cio.PrintMessage(dstImgI1B)
}
