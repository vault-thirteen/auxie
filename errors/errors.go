package errors

import (
	"errors"
)

const ErrorsSeparator = "; "

// Combine combines two errors into a single error.
func Combine(error1 error, error2 error) error {
	if error1 == nil {
		if error2 == nil {
			return nil
		}
		return error2
	}

	if error2 == nil {
		return error1
	}

	return errors.New(error1.Error() + ErrorsSeparator + error2.Error())
}
