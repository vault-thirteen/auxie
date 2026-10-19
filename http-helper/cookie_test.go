package httphelper

import (
	"net/http"
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_GetCookieByName(t *testing.T) {
	aTest := tester.New(t)
	var req *http.Request
	var c1, c2 *http.Cookie
	var err error

	// Test #1. OK.
	req = &http.Request{}
	c1 = &http.Cookie{
		Name:  "abc",
		Value: "123",
	}
	req.Header = make(http.Header)
	req.AddCookie(c1)
	c2, err = GetCookieByName(req, "abc")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(c2, c1)

	// Test #2. No cookies.
	req = &http.Request{}
	req.Header = make(http.Header)
	c2, err = GetCookieByName(req, "abc")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(c2, (*http.Cookie)(nil))

	// Test #3. Duplicate cookies.
	req = &http.Request{}
	c1 = &http.Cookie{
		Name:  "abc",
		Value: "123",
	}
	req.Header = make(http.Header)
	req.AddCookie(c1)
	req.AddCookie(c1)
	c2, err = GetCookieByName(req, "abc")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(c2, (*http.Cookie)(nil))
}
