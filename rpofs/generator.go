package rpofs

import (
	"errors"

	"github.com/vault-thirteen/auxie/random"
)

type Generator struct {
	passwordLength int
	allowedSymbols []rune
}

func NewGenerator(pwdSize int, allowedSymbols []rune) (g *Generator, err error) {
	if pwdSize < 1 {
		return nil, errors.New(ErrPasswordLengthTooSmall)
	}

	if pwdSize > PasswordLengthMax {
		return nil, errors.New(ErrPasswordLengthTooBig)
	}

	if len(allowedSymbols) < 1 {
		return nil, errors.New(ErrAllowedPasswordSymbolsSetShort)
	}

	if !isSetOfUniqueSymbols(allowedSymbols) {
		return nil, errors.New(ErrAllowedPasswordSymbolsSetNotUnique)
	}

	g = &Generator{
		passwordLength: pwdSize,
		allowedSymbols: allowedSymbols,
	}

	return
}

func (g *Generator) CreatePassword() (password *string, err error) {
	buffer := make([]rune, g.passwordLength)
	sliceIndexMax := uint(len(g.allowedSymbols) - 1)
	var allowedSymbolIdx uint

	for bufferCursor := 0; bufferCursor < g.passwordLength; bufferCursor++ {
		allowedSymbolIdx, err = random.Uint(0, sliceIndexMax)
		if err != nil {
			return
		}

		buffer[bufferCursor] = g.allowedSymbols[allowedSymbolIdx]
	}

	bufferString := string(buffer)

	return &bufferString, nil
}
