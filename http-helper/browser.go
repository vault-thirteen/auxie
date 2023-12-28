package httphelper

import (
	"net/http"

	mime "github.com/vault-thirteen/auxie/MIME"
	"github.com/vault-thirteen/auxie/header"
)

func CheckBrowserSupportForJson(req *http.Request) (ok bool, err error) {
	var httpAcceptHeader string
	httpAcceptHeader, err = GetSingleHttpHeader(req, header.HttpHeaderAccept)
	if err != nil {
		return false, err
	}

	var amts *AcceptedMimeTypes
	amts, err = NewAcceptedMimeTypesFromHeader(httpAcceptHeader)
	if err != nil {
		return false, err
	}

	var amt *AcceptedMimeType
	for {
		amt, err = amts.Next()
		if err != nil {
			break
		}

		switch amt.MimeType {
		case mime.TypeApplicationJson,
			mime.TypeApplicationAny,
			mime.TypeAny:
			return true, nil
		}
	}

	return false, nil
}
