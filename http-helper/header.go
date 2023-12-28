package httphelper

import (
	"errors"
	"fmt"
	"net/http"
	"net/textproto"
	"strings"
)

// Errors
const (
	ErrHttpHeaderNameIsNotSet   = "HTTP header name is not set"
	ErrHttpHeaderNameIsNotFound = "HTTP header name is not found"
	ErrFMultipleHttpHeaders     = "multiple HTTP headers: %s"
	ErrFHttpHeaderIsNotFound    = "HTTP header is not found: %s"
)

// FindHttpHeader function tries to find the HTTP header with name similar to
// the specified one. According to the Section 4.2 of RFC 2616, HTTP header
// names are case-insensitive. On success, it returns the exact header name
// which was found.
//
// Notes. This function is useful for inspection of HTTP requests having
// incorrect HTTP headers. It tries to show an actual name of the incorrectly
// set header. For example, if a request has a header named 'CONtent-TyPe',
// this function will search for headers similar to 'Content-Type' and find the
// modified variant of it and return it as it is, i.e. as 'CONtent-TyPe'.
// Most users do not need this function.
func FindHttpHeader(req *http.Request, headerNameAsked string) (hdrNm string, err error) {
	var headerExists bool

	if req == nil {
		return hdrNm, errors.New(ErrNullPointer)
	}
	if len(headerNameAsked) == 0 {
		return hdrNm, errors.New(ErrHttpHeaderNameIsNotSet)
	}

	// 1. Try the easy way.
	_, headerExists = req.Header[headerNameAsked]
	if headerExists {
		return headerNameAsked, nil
	}

	// 2. Try the canonical header name.
	headerNameAskedCN := textproto.CanonicalMIMEHeaderKey(headerNameAsked)
	_, headerExists = req.Header[headerNameAskedCN]
	if headerExists {
		return headerNameAskedCN, nil
	}

	// 3. Try the difficult way.
	headerNameAskedLC := strings.ToLower(headerNameAsked)
	for headerName := range req.Header {
		if strings.ToLower(headerName) == headerNameAskedLC {
			return headerName, nil
		}
	}

	return hdrNm, errors.New(ErrHttpHeaderNameIsNotFound)
}

// DeleteHttpHeader function tries to delete a header from the HTTP request.
func DeleteHttpHeader(r *http.Request, headerNameToDelete string) (err error) {
	var headerNameExact string

	// Find the header.
	headerNameExact, err = FindHttpHeader(r, headerNameToDelete)
	if err != nil {
		return err
	}

	// Delete the header.
	delete(r.Header, headerNameExact)

	return nil
}

// GetSingleHttpHeader reads exactly one, single HTTP header. If multiple
// headers are found, an error is returned. Unfortunately, Go language does not
// do this by default and returns only a first header value even when there are
// multiple values available. Such a behaviour may lead to unexpected errors.
func GetSingleHttpHeader(req *http.Request, headerName string) (h string, err error) {
	headerNameCN := textproto.CanonicalMIMEHeaderKey(headerName)

	headers := req.Header[headerNameCN]

	if len(headers) > 1 {
		return "", fmt.Errorf(ErrFMultipleHttpHeaders, headerNameCN)
	}

	if len(headers) == 0 {
		return "", fmt.Errorf(ErrFHttpHeaderIsNotFound, headerNameCN)
	}

	return headers[0], nil
}
