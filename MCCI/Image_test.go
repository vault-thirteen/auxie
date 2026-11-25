package mcci

import (
	"image"
	"os"
	"path/filepath"
	"testing"

	"github.com/vault-thirteen/auxie/cio"
	"github.com/vault-thirteen/auxie/colour"
	e "github.com/vault-thirteen/auxie/os/win32/enum"
	"github.com/vault-thirteen/auxie/tester"

	_ "image/gif"  // For the GIF format decoder.
	_ "image/jpeg" // For the JPEG format decoder.
	_ "image/png"  // For the PNG format decoder.

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
)

const (
	TestFolder = "test"
)

func Test_NewImageFromImage_NRGBA(t *testing.T) {
	aTest := tester.New(t)
	filePath := filepath.Join(TestFolder, "3x2.png")

	f, err := os.Open(filePath)
	aTest.MustBeNoError(err)
	defer func() {
		derr := f.Close()
		aTest.MustBeNoError(derr)
	}()

	var srcImg image.Image
	var formatName string
	srcImg, formatName, err = image.Decode(f)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(formatName, "png")
	aTest.MustBeEqual(c.GetColourModelType(srcImg.ColorModel()), c.ColourModelType_NRGBA)

	var x any
	var ok bool
	var dstImgI1B *Image[uint8]
	var dstImgI2B *Image[uint16]
	var dstImgF4B *Image[float32]

	// Straight.
	{
		cio.PrintMessage("Straight")

		x, err = NewImageFromImage(srcImg, PixelFormat_Integer_1B, AlphaChannelMode_Straight)
		aTest.MustBeNoError(err)
		dstImgI1B, ok = x.(*Image[uint8])
		aTest.MustBeEqual(ok, true)
		cio.PrintMessage(dstImgI1B)

		x, err = NewImageFromImage(srcImg, PixelFormat_Integer_2B, AlphaChannelMode_Straight)
		aTest.MustBeNoError(err)
		dstImgI2B, ok = x.(*Image[uint16])
		aTest.MustBeEqual(ok, true)
		cio.PrintMessage(dstImgI2B)

		x, err = NewImageFromImage(srcImg, PixelFormat_Float_4B, AlphaChannelMode_Straight)
		aTest.MustBeNoError(err)
		dstImgF4B, ok = x.(*Image[float32])
		aTest.MustBeEqual(ok, true)
		cio.PrintMessage(dstImgF4B)
	}

	// Pre-multiplied.
	{
		cio.PrintMessage("Pre-multiplied")

		x, err = NewImageFromImage(srcImg, PixelFormat_Integer_1B, AlphaChannelMode_Premultiplied)
		aTest.MustBeNoError(err)
		dstImgI1B, ok = x.(*Image[uint8])
		aTest.MustBeEqual(ok, true)
		cio.PrintMessage(dstImgI1B)

		x, err = NewImageFromImage(srcImg, PixelFormat_Integer_2B, AlphaChannelMode_Premultiplied)
		aTest.MustBeNoError(err)
		dstImgI2B, ok = x.(*Image[uint16])
		aTest.MustBeEqual(ok, true)
		cio.PrintMessage(dstImgI2B)

		x, err = NewImageFromImage(srcImg, PixelFormat_Float_4B, AlphaChannelMode_Premultiplied)
		aTest.MustBeNoError(err)
		dstImgF4B, ok = x.(*Image[float32])
		aTest.MustBeEqual(ok, true)
		cio.PrintMessage(dstImgF4B)
	}
}
func Test_ExportToDirect2D(t *testing.T) {
	aTest := tester.New(t)
	filePath := filepath.Join(TestFolder, "3x2.png")

	f, err := os.Open(filePath)
	aTest.MustBeNoError(err)
	defer func() {
		derr := f.Close()
		aTest.MustBeNoError(derr)
	}()

	var srcImg image.Image
	srcImg, _, err = image.Decode(f)
	aTest.MustBeNoError(err)

	var x any
	x, err = NewImageFromImage(srcImg, PixelFormat_Integer_1B, AlphaChannelMode_Straight)
	aTest.MustBeNoError(err)
	var img1B *Image[uint8]
	var ok bool
	img1B, ok = x.(*Image[uint8])
	aTest.MustBeEqual(ok, true)

	var buf []byte
	buf, err = img1B.ExportToDirect2D(e.DXGI_FORMAT_B8G8R8A8_UNORM, e.D2D1_ALPHA_MODE_PREMULTIPLIED)
	aTest.MustBeNoError(err)
	cio.PrintMessage(buf)
}
