// Reader_test.go.

package reader

import (
	"bytes"
	"io"
	"testing"
)

func Test_NewReader(t *testing.T) {

	var reader io.Reader
	var result *Reader

	reader = bytes.NewReader([]byte{})
	result = NewReader(reader)
	if result.reader != reader {
		t.FailNow()
	}
}
