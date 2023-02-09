// utf.go.

// Unicode Symbols Processing.

package unicode

import (
	"math"
	"unicode/utf8"
)

// CreateValidUtf8Runes Function creates a full Set of available Unicode UTF-8
// Symbols.
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
