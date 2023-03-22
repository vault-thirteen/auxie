package unicode

import (
	"os"
	"testing"
	"time"
	"unicode/utf8"
)

func Test_CreateUtf8Runes(t *testing.T) {
	const tmpFileName = "runes.tmp.txt"

	var err error
	var file *os.File
	var runes []rune

	file, err = os.OpenFile(tmpFileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		t.FailNow()
	}

	runes = CreateValidUtf8Runes()
	for _, aRune := range runes {
		if !utf8.ValidRune(aRune) {
			t.FailNow()
		}
		_, err = file.Write([]byte(string(aRune)))
		if err != nil {
			t.FailNow()
		}
	}
	if len(runes) != 1112064 {
		t.FailNow()
	}

	err = file.Close()
	if err != nil {
		t.FailNow()
	}

	time.Sleep(time.Second * 1) // For manual check.

	err = os.Remove(tmpFileName)
	if err != nil {
		t.FailNow()
	}
}
