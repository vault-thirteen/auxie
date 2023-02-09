// symbol_test.go.

package unicode

import (
	"testing"
)

func Test_symbolIsRusLatLetter(t *testing.T) {

	// Test #1. Allowed Symbols.
	lettsersStr := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"абвгдеёжзиклмнопрстуфхцчшщъыьэюя" +
		"АБВГДЕЁЖЗИКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"
	letters := []rune(lettsersStr)
	for _, letter := range letters {
		if SymbolIsRusLatLetter(letter) != true {
			t.FailNow()
		}
	}

	// Test #2. Bad Symbol.
	if SymbolIsRusLatLetter('9') != false {
		t.FailNow()
	}
}

func Test_symbolIsLatLetter(t *testing.T) {

	// Test #1. Allowed Symbols.
	lettsersStr := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letters := []rune(lettsersStr)
	for _, letter := range letters {
		if SymbolIsLatLetter(letter) != true {
			t.FailNow()
		}
	}

	// Test #2. Bad Symbol.
	if SymbolIsLatLetter('Ы') != false {
		t.FailNow()
	}
}

func Test_symbolIsRusLetter(t *testing.T) {

	// Test #1. Allowed Symbols.
	lettsersStr := "абвгдеёжзиклмнопрстуфхцчшщъыьэюя" +
		"АБВГДЕЁЖЗИКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"
	letters := []rune(lettsersStr)
	for _, letter := range letters {
		if SymbolIsRusLetter(letter) != true {
			t.FailNow()
		}
	}

	// Test #2. Bad Symbol.
	if SymbolIsRusLetter('X') != false {
		t.FailNow()
	}
}

func Test_symbolIsNumber(t *testing.T) {

	// Test #1. Allowed Symbols.
	symbolsStr := "0123456789"
	symbols := []rune(symbolsStr)
	for _, symbol := range symbols {
		if SymbolIsNumber(symbol) != true {
			t.FailNow()
		}
	}

	// Test #2. Bad Symbol.
	if SymbolIsNumber('X') != false {
		t.FailNow()
	}
}
