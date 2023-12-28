package httphelper

import (
	"net/http"
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_CheckBrowserSupportForJson(t *testing.T) {
	var aTest = tester.New(t)
	var ok bool
	var err error
	var request *http.Request

	// Test #1. Multiple headers.
	request = &http.Request{
		Header: map[string][]string{
			"Accept": []string{"1", "2", "3"},
		},
	}
	ok, err = CheckBrowserSupportForJson(request)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(ok, false)

	// Test #2. No headers.
	request = &http.Request{
		Header: map[string][]string{
			"X-X": {"(o.0)"},
		},
	}
	ok, err = CheckBrowserSupportForJson(request)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(ok, false)

	// Test #3. Single header. Yes.
	request = &http.Request{
		Header: map[string][]string{
			"Accept": []string{"*/*"},
		},
	}
	ok, err = CheckBrowserSupportForJson(request)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(ok, true)

	// Test #4. Single header. No.
	request = &http.Request{
		Header: map[string][]string{
			"Accept": []string{"text/plain"},
		},
	}
	ok, err = CheckBrowserSupportForJson(request)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(ok, false)

	// Test #5. Single header. Broken syntax.
	request = &http.Request{
		Header: map[string][]string{
			"Accept": []string{"text/plain;;"},
		},
	}
	ok, err = CheckBrowserSupportForJson(request)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(ok, false)
}
