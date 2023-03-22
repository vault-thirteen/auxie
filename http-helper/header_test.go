package httphelper

import (
	"net/http"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_FindHTTPHeader(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var headerName string
	var request *http.Request

	// Test #1. Null Request.
	request = nil
	headerName, err = FindHTTPHeader(request, "abc")
	aTest.MustBeAnError(err)

	// Test #2. Empty Header Name.
	request = &http.Request{}
	headerName, err = FindHTTPHeader(request, "")
	aTest.MustBeAnError(err)

	// Test #3. Existent Header with exact Name Match.
	request = &http.Request{
		Header: map[string][]string{
			"Content-Type": {"Intergalactic Message"},
		},
	}
	headerName, err = FindHTTPHeader(request, "Content-Type")
	aTest.MustBeNoError(err)
	if headerName != "Content-Type" {
		t.FailNow()
	}

	// Test #4. Existent Header with similar Name.
	request = &http.Request{
		Header: map[string][]string{
			"Content-Type": {"Intergalactic Message"},
		},
	}
	headerName, err = FindHTTPHeader(request, "content-type")
	aTest.MustBeNoError(err)
	if headerName != "Content-Type" {
		t.FailNow()
	}

	// Test #5. Non-existent Header.
	request = &http.Request{
		Header: map[string][]string{
			"Content-Type": {"Intergalactic Message"},
		},
	}
	headerName, err = FindHTTPHeader(request, "X-FakeHeader")
	aTest.MustBeAnError(err)
}

func Test_DeleteHTTPHeader(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var request *http.Request

	// Test #1. Test of Entry into the 'FindHTTPHeader' Function.
	request = nil
	err = DeleteHTTPHeader(request, "abc")
	aTest.MustBeAnError(err)

	// Test #2. Test of real Deletion.
	request = &http.Request{
		Header: map[string][]string{
			"Content-Type": {"Intergalactic Message"},
			"X-Service":    {"Intergalactic Service"},
		},
	}
	err = DeleteHTTPHeader(request, "x-service")
	aTest.MustBeNoError(err)
	if (len(request.Header) != 1) ||
		(len(request.Header["Content-Type"]) != 1) ||
		(request.Header["Content-Type"][0] != "Intergalactic Message") {
		t.FailNow()
	}
}
