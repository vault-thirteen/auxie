package unicode

// Unicode Symbols Processing.

import (
	"math"
	"unicode/utf8"
)

// CreateValidUtf8Runes creates a full set of available UTF-8 Unicode symbols.
func CreateValidUtf8Runes() []rune {
	var r rune
	var runes = make([]rune, 0)

	r = 0
	for {
		if utf8.ValidRune(r) {
			runes = append(runes, r)
		}

		// Next Rune.
		if r == math.MaxInt32 {
			break
		}
		r++
	}

	return runes
}
