package httphelper

import (
	"net/http"
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_FindHttpHeader(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var headerName string
	var request *http.Request

	// Test #1. Null request.
	request = nil
	headerName, err = FindHttpHeader(request, "abc")
	aTest.MustBeAnError(err)

	// Test #2. Empty header name.
	request = &http.Request{}
	headerName, err = FindHttpHeader(request, "")
	aTest.MustBeAnError(err)

	// Test #3. Difficult way.
	request = &http.Request{
		Header: map[string][]string{
			"CoNtEnT-TyPe": {"Vandalised header"},
		},
	}
	headerName, err = FindHttpHeader(request, "content-type")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(headerName, "CoNtEnT-TyPe")

	// Test #4. Canonical way.
	request = &http.Request{
		Header: map[string][]string{
			"Content-Type": {"Intergalactic Message"},
		},
	}
	headerName, err = FindHttpHeader(request, "content-type")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(headerName, "Content-Type")

	// Test #5. Easy way.
	request = &http.Request{
		Header: map[string][]string{
			"content-type": {"our keyboard is broken, no capital letters"},
		},
	}
	headerName, err = FindHttpHeader(request, "content-type")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(headerName, "content-type")

	// Test #6. Header is not found.
	request = &http.Request{
		Header: map[string][]string{
			"X-FakeHeader": {"boo"},
		},
	}
	headerName, err = FindHttpHeader(request, "content-type")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(err.Error(), ErrHttpHeaderNameIsNotFound)
}

func Test_DeleteHttpHeader(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var request *http.Request

	// Test #1. Test of entry into the 'FindHttpHeader' function.
	request = nil
	err = DeleteHttpHeader(request, "abc")
	aTest.MustBeAnError(err)

	// Test #2. Test of real deletion.
	request = &http.Request{
		Header: map[string][]string{
			"Content-Type": []string{"1", "2", "3"},
			"X-SeRvIcE":    {"JuNkY tOwN"},
		},
	}
	err = DeleteHttpHeader(request, "x-service")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(len(request.Header), 1)
	aTest.MustBeEqual(len(request.Header["Content-Type"]), 3)
	aTest.MustBeEqual(request.Header["Content-Type"][0], "1")
}

func Test_GetSingleHttpHeader(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var headerName string
	var request *http.Request

	// Test #1. Multiple headers.
	request = &http.Request{
		Header: map[string][]string{
			"Content-Type": []string{"1", "2", "3"},
			"X-X":          {"(o.0)"},
		},
	}
	headerName, err = GetSingleHttpHeader(request, "content-type")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(headerName, "")

	// Test #2. No headers.
	request = &http.Request{
		Header: map[string][]string{
			"X-X": {"(o.0)"},
		},
	}
	headerName, err = GetSingleHttpHeader(request, "content-type")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(headerName, "")

	// Test #3. Single header.
	request = &http.Request{
		Header: map[string][]string{
			"Hooeelo": {"0783 1505"},
			"X-X":     {"(o.0)"},
		},
	}
	headerName, err = GetSingleHttpHeader(request, "x-x")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(headerName, "(o.0)")
}
