// generator.go.

package rpofs

import "github.com/vault-thirteen/auxie/random"

type Generator struct {
	passwordLength int
	allowedSymbols []rune
}

func NewGenerator(
	passwordFixedSize int,
	allowedPasswordSymbols []rune,
) (generator *Generator, err error) {
	if passwordFixedSize < 1 {
		return nil, errPasswordLengthTooSmall
	}

	if passwordFixedSize > PasswordLengthMax {
		return nil, errPasswordLengthTooBig
	}

	if len(allowedPasswordSymbols) < 1 {
		return nil, errAllowedPasswordSymbolsSetShort
	}

	if !isSetOfUniqueSymbols(allowedPasswordSymbols) {
		return nil, errAllowedPasswordSymbolsSetNotUnique
	}

	generator = &Generator{
		passwordLength: passwordFixedSize,
		allowedSymbols: allowedPasswordSymbols,
	}

	return
}

func (g *Generator) CreatePassword() (
	password *string,
	err error,
) {
	buffer := make([]rune, g.passwordLength)
	sliceIndexMax := len(g.allowedSymbols) - 1

	var allowedSymbolIdx uint

	for bufferCursor := 0; bufferCursor < g.passwordLength; bufferCursor++ {
		allowedSymbolIdx, err = random.Uint(0, uint(sliceIndexMax))
		if err != nil {
			return
		}

		buffer[bufferCursor] = g.allowedSymbols[allowedSymbolIdx]
	}

	bufferString := string(buffer)

	return &bufferString, nil
}
