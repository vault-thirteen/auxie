package httphelper

import (
	"net/http"

	"github.com/vault-thirteen/auxie/errors"
)

const (
	Errf_DuplicateCookie = "duplicate cookie: %s"
)

// GetCookieByName reads a non-duplicate cookie. If a cookie is not found, null
// is returned.
func GetCookieByName(req *http.Request, cookieName string) (cookie *http.Cookie, err error) {
	cookiesFound := make(map[string]*http.Cookie)

	allCookies := req.Cookies()

	for _, cookie = range allCookies {
		if cookie.Name == cookieName {
			_, alreadyExists := cookiesFound[cookieName]
			if alreadyExists {
				return nil, errors.ErrorFromFormattedText(Errf_DuplicateCookie, cookieName)
			}

			cookiesFound[cookieName] = cookie
		}
	}

	return cookiesFound[cookieName], nil
}
