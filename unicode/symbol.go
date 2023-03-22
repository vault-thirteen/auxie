package unicode

// SymbolIsRusLatLetter checks whether the specified symbol is a letter of
// Russian or Latin alphabets or not.
func SymbolIsRusLatLetter(symbol rune) bool {
	if SymbolIsLatLetter(symbol) {
		return true
	}
	if SymbolIsRusLetter(symbol) {
		return true
	}
	return false
}

// SymbolIsLatLetter checks whether the specified symbol is a letter of
// Latin alphabet or not.
func SymbolIsLatLetter(symbol rune) bool {
	if (symbol >= 'a') && (symbol <= 'z') {
		return true
	}
	if (symbol >= 'A') && (symbol <= 'Z') {
		return true
	}
	return false
}

// SymbolIsRusLetter checks whether the specified symbol is a letter of
// Russian alphabet or not.
func SymbolIsRusLetter(symbol rune) bool {
	if (symbol >= 'а') && (symbol <= 'я') {
		return true
	}
	if (symbol >= 'А') && (symbol <= 'Я') {
		return true
	}
	if symbol == 'ё' {
		return true
	}
	if symbol == 'Ё' {
		return true
	}
	return false
}

// SymbolIsNumber checks whether the specified symbol is numeric.
func SymbolIsNumber(symbol rune) bool {
	if (symbol >= '0') && (symbol <= '9') {
		return true
	}
	return false
}
