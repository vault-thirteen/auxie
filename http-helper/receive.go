package httphelper

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"

	"github.com/vault-thirteen/errorz"
)

// Functions which help in receiving Data from HTTP requests.

// ReceiveJSON function receives an object encoded with JSON format from the
// HTTP request's body.
func ReceiveJSON(r *http.Request, receiver interface{}) (err error) {
	var bodyContents []byte
	var jsonDecoder *json.Decoder

	if reflect.TypeOf(receiver).Kind() != reflect.Ptr {
		return errors.New(ErrNotPointer)
	}

	bodyContents, err = io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer func() {
		derr := r.Body.Close()
		err = errorz.Combine(err, derr)
	}()

	jsonDecoder = json.NewDecoder(bytes.NewReader(bodyContents))
	err = jsonDecoder.Decode(receiver)
	if err != nil {
		return err
	}

	return nil
}
