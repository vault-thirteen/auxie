package m

import "unicode"

func CapitalizeFirstLetter(in string) (out string) {
	if len(in) == 0 {
		return ""
	}

	runes := []rune(in)
	if len(runes) == 0 {
		return ""
	}

	runes[0] = unicode.ToUpper(runes[0])

	return string(runes)
}
