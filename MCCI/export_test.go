package mcci

import (
	"image"
	"os"
	"path/filepath"
	"testing"

	"github.com/vault-thirteen/auxie/cio"
	e "github.com/vault-thirteen/auxie/os/win32/enum"
	"github.com/vault-thirteen/auxie/tester"
)

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
