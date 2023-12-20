package httphelper

import (
	"net/http"
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_PerformSimpleHttpTest(t *testing.T) {

	var aSimpleHttpTest SimpleTest
	var aTest = tester.New(t)
	var err error

	// Test #1. Positive.
	aSimpleHttpTest = SimpleTest{
		Parameter: SimpleTestParameter{
			RequestMethod:  "GET",
			RequestUrl:     "http://example.org",
			RequestBody:    nil,
			RequestHandler: nil, // Is set below.
		},
		ResultExpected: SimpleTestResult{
			ResponseBodyString: "Hello",
			ResponseStatusCode: http.StatusAccepted,
		},
	}
	aSimpleHttpTest.Parameter.RequestHandler = func(w http.ResponseWriter, r *http.Request) {
		var herr error
		w.WriteHeader(http.StatusAccepted)
		_, herr = w.Write([]byte("Hello"))
		if herr != nil {
			t.FailNow()
		}
	}
	err = PerformSimpleHttpTest(&aSimpleHttpTest)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(aSimpleHttpTest.ResultReceived, aSimpleHttpTest.ResultExpected)

	// Test #2. Negative.
	aSimpleHttpTest = SimpleTest{
		Parameter: SimpleTestParameter{
			RequestMethod:  "GET",
			RequestUrl:     "http://example.org",
			RequestBody:    nil,
			RequestHandler: nil, // Is set below.
		},
		ResultExpected: SimpleTestResult{
			ResponseBodyString: "Hello",
			ResponseStatusCode: http.StatusAccepted,
		},
	}
	aSimpleHttpTest.Parameter.RequestHandler = func(w http.ResponseWriter, r *http.Request) {
		var herr error
		w.WriteHeader(http.StatusBadRequest)
		_, herr = w.Write([]byte("Bye-bye"))
		if herr != nil {
			t.FailNow()
		}
	}
	err = PerformSimpleHttpTest(&aSimpleHttpTest)
	aTest.MustBeNoError(err)
	aTest.MustBeDifferent(aSimpleHttpTest.ResultReceived, aSimpleHttpTest.ResultExpected)
}
