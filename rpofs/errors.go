// errors.go.

package rpofs

import "errors"

var (
	errPasswordLengthTooSmall = errors.New("password's length is " +
		"too small")
	errPasswordLengthTooBig = errors.New("password's length is " +
		"too big")
	errAllowedPasswordSymbolsSetShort = errors.New("allowed password " +
		"symbols set is too short")
	errAllowedPasswordSymbolsSetNotUnique = errors.New("allowed " +
		"password symbols set is not unique")
)
