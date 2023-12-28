package httphelper

import (
	"encoding/json"
	"net/http"
)

// Functions which help in receiving Data from HTTP requests.

// ReceiveJSON function receives an object encoded with JSON format from the
// HTTP request's body. The receiver must be a pointer.
func ReceiveJSON(r *http.Request, receiver interface{}) (err error) {
	return json.NewDecoder(r.Body).Decode(receiver)
}
