package httphelper

import (
	"encoding/json"
	"net/http"

	"github.com/vault-thirteen/auxie/MIME"
	"github.com/vault-thirteen/auxie/header"
)

// Errors.
const (
	ErrNullPointer = "null pointer"
)

// Functions which help in replying to HTTP requests.

// ReplyTextWithCode function replies to the HTTP request with the specified
// text and HTTP status code.
func ReplyTextWithCode(rw http.ResponseWriter, httpStatusCode int, replyText string) {
	rw.WriteHeader(httpStatusCode)

	_, err := rw.Write([]byte(replyText))
	logErrorIfSet(err)
}

// ReplyErrorWithCode function replies to the HTTP request with an error and
// the specified HTTP status code.
func ReplyErrorWithCode(rw http.ResponseWriter, httpStatusCode int, err error) {
	ReplyTextWithCode(rw, httpStatusCode, err.Error())
}

// ReplyErrorInternal function replies to the HTTP request with an error and
// 'Internal Server Error' HTTP status code.
func ReplyErrorInternal(rw http.ResponseWriter, err error) {
	ReplyErrorWithCode(rw, http.StatusInternalServerError, err)
}

// ReplyJSON function sends an object in JSON format to the HTTP output stream.
func ReplyJSON(rw http.ResponseWriter, replyObject interface{}) {
	// Encode an object with JSON format.
	response, err := json.Marshal(replyObject)
	logErrorIfSet(err)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}

	// Send the reply.
	rw.Header().Set(header.HttpHeaderContentType, mime.TypeApplicationJson)

	_, err = rw.Write(response)
	logErrorIfSet(err)
}

// ReplyJSONFast function sends an object in JSON format to the HTTP output
// stream using the faster but less secure way than an ordinary 'ReplyJSON'
// method.
func ReplyJSONFast(rw http.ResponseWriter, replyObject interface{}) {
	// Encode an object with JSON format and send it simultaneously.
	rw.Header().Set(header.HttpHeaderContentType, mime.TypeApplicationJson)

	err := json.NewEncoder(rw).Encode(replyObject)
	logErrorIfSet(err)
}
