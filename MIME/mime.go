package mime

import (
	"slices"
	"strings"

	"github.com/vault-thirteen/auxie/errors"
)

// Common MIME Types.

const (
	Errf_Syntax          = "syntax error in MIME type: %s"
	Errf_UnknownCategory = "unknown category: %s"
)

const (
	TypeAny = "*/*"
)

var categories = []string{
	"*",
	"application",
	"audio",
	"font",
	"image",
	"model",
	"multipart",
	"text",
	"video",
}

func GetCategories() []string {
	return categories
}

func IsCategoryValid(cat string) (ok bool) {
	return slices.Contains(categories, cat)
}

// Unfortunately, although Go language supports some kind of reflection,
// including functionality to get name of types and struct fields, it is
// unable to list all variables and constants of a package. So, we can not
// create a function to check whether a specified string is a valid MIME type
// or not. Maybe, in a hundred of years they will add this functionality, but
// this language will be useless at that time for sure.

func GetMimeTypeParts(s string) (category string, subtype string, err error) {
	parts := strings.Split(s, "/")

	if len(parts) != 2 {
		return "", "", errors.ErrorFromFormattedText(Errf_Syntax, s)
	}

	if !IsCategoryValid(parts[0]) {
		return "", "", errors.ErrorFromFormattedText(Errf_UnknownCategory, parts[0])
	}

	if parts[0] == `*` {
		if parts[1] != `*` {
			return "", "", errors.ErrorFromFormattedText(Errf_Syntax, s)
		}
	}

	// We can not check the sub-type at this moment.

	return parts[0], parts[1], nil
}
