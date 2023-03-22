package httphelper

import (
	"errors"
	"net/http"
	"strings"
)

// Errors
const (
	ErrHTTPHeaderNameIsNotSet   = "HTTP Header Name is not set"
	ErrHTTPHeaderNameIsNotFound = "HTTP Header Name is not found"
)

// FindHTTPHeader function tries to find the HTTP header with name similar to
// the specified one. According to the Section 4.2 of RFC 2616, HTTP header
// names are case-insensitive. On success, returns 'nil' and the exact header
// name which was found.
func FindHTTPHeader(r *http.Request, headerNameAsked string) (hdr string, err error) {
	var headerExists bool

	if r == nil {
		return hdr, errors.New(ErrNullPointer)
	}
	if len(headerNameAsked) == 0 {
		return hdr, errors.New(ErrHTTPHeaderNameIsNotSet)
	}

	// 1. Try the easy Way.
	_, headerExists = r.Header[headerNameAsked]
	if headerExists {
		return headerNameAsked, nil
	}

	// 2. Try the difficult Way.
	headerNameAskedLC := strings.ToLower(headerNameAsked)
	for headerName := range r.Header {
		headerNameLC := strings.ToLower(headerName)
		if headerNameLC == headerNameAskedLC {
			return headerName, nil
		}
	}

	return hdr, errors.New(ErrHTTPHeaderNameIsNotFound)
}

// DeleteHTTPHeader function tries to delete a header from the HTTP request.
func DeleteHTTPHeader(r *http.Request, headerNameToDelete string) (err error) {
	var headerNameExact string

	// Find the Header.
	headerNameExact, err = FindHTTPHeader(r, headerNameToDelete)
	if err != nil {
		return err
	}

	// Delete the Header.
	delete(r.Header, headerNameExact)
	return nil
}
