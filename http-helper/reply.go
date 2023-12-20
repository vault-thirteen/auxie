package httphelper

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/vault-thirteen/MIME"
	"github.com/vault-thirteen/auxie/header"
)

// Errors.
const (
	ErrNullPointer = "null pointer"
)

// Functions which help in replying to HTTP requests.

// ReplyTextWithCode function replies to the HTTP request with the specified
// text and HTTP status code.
func ReplyTextWithCode(w http.ResponseWriter, httpStatusCode int, replyText string) {
	w.WriteHeader(httpStatusCode)
	_, xerr := w.Write([]byte(replyText))
	if xerr != nil {
		log.Println(replyText)
		log.Println(xerr)
	}
}

// ReplyErrorWithCode function replies to the HTTP request with an error and
// the specified HTTP status code.
func ReplyErrorWithCode(w http.ResponseWriter, httpStatusCode int, err error) {
	ReplyTextWithCode(w, httpStatusCode, err.Error())
}

// ReplyErrorInternal function replies to the HTTP request with an error and
// 'Internal Server Error' HTTP status code.
func ReplyErrorInternal(w http.ResponseWriter, err error) {
	ReplyErrorWithCode(w, http.StatusInternalServerError, err)
}

// ReplyJSON function sends an object in JSON format to the HTTP output stream.
func ReplyJSON(w http.ResponseWriter, replyObject interface{}) {
	var err error
	var response []byte

	// Encode an Object with JSON Format.
	response, err = json.Marshal(replyObject)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}

	// Send the Reply.
	w.Header().Set(header.HttpHeaderContentType, mime.TypeApplicationJson)
	_, err = w.Write(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}

// ReplyJSONFast function sends an object in JSON format to the HTTP output
// stream using the faster but less secure way than an ordinary 'ReplyJSON'
// method.
func ReplyJSONFast(w http.ResponseWriter, replyObject interface{}) {
	var err error
	var jsonEncoder *json.Encoder

	// Create the JSON Encoder.
	jsonEncoder = json.NewEncoder(w)
	if jsonEncoder == nil {
		err = errors.New(ErrNullPointer)
		log.Println(err)
	}

	// Encode an Object with JSON Format and send it simultaneously.
	w.Header().Set(header.HttpHeaderContentType, mime.TypeApplicationJson)
	err = jsonEncoder.Encode(replyObject)
	if err != nil {
		log.Println(err)
	}
}
