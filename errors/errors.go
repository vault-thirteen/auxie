package errors

import (
	"errors"
	"fmt"
	"log"
	"strings"
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

// ErrorFromText creates an error from text.
func ErrorFromText(text string) error {
	tt := strings.TrimSpace(text)

	if len(tt) == 0 {
		return nil
	}

	return errors.New(tt)
}

// ErrorFromFormattedText creates an error from formatted text.
func ErrorFromFormattedText(format string, args ...any) error {
	tf := strings.TrimSpace(format)

	if len(tf) == 0 {
		return nil
	}

	return fmt.Errorf(tf, args...)
}

// ErrorWithPrefix wraps an existing error with specified textual prefix.
// Textual prefix must not contain colon at the end, as it is appended
// automatically.
func ErrorWithPrefix(err error, prefix string) error {
	if err == nil {
		return nil
	}

	tp := strings.TrimSpace(prefix)

	if len(tp) == 0 {
		return err
	}

	return fmt.Errorf(tp+": %w", err)
}

// LogError simply logs the error.
func LogError(err error) {
	log.Println(err)
}
