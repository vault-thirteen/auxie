// reply_test.go.

package httphelper

import (
	"errors"
	"net/http"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_ReplyTextWithCode(t *testing.T) {

	var err error
	var httpTest SimpleTest
	var test *tester.Test

	test = tester.New(t)
	httpTest = SimpleTest{
		Parameter: SimpleTestParameter{
			RequestMethod:  TestMethod,
			RequestUrl:     TestUrl,
			RequestBody:    nil,
			RequestHandler: nil, // Is set below.
		},
		ResultExpected: SimpleTestResult{
			ResponseBodyString: "Some Text",
			ResponseStatusCode: http.StatusBadRequest,
		},
	}
	httpTest.Parameter.RequestHandler = func(w http.ResponseWriter, r *http.Request) {
		ReplyTextWithCode( // <- This HTTP Handler Function is being tested.
			w,
			httpTest.ResultExpected.ResponseStatusCode,
			httpTest.ResultExpected.ResponseBodyString,
		)
	}
	err = PerformSimpleHttpTest(&httpTest)
	test.MustBeNoError(err)
	test.MustBeEqual(httpTest.ResultReceived, httpTest.ResultExpected)
}

func Test_ReplyErrorWithCode(t *testing.T) {

	var err error
	var httpTest SimpleTest
	var test *tester.Test

	test = tester.New(t)
	httpTest = SimpleTest{
		Parameter: SimpleTestParameter{
			RequestMethod:  TestMethod,
			RequestUrl:     TestUrl,
			RequestBody:    nil,
			RequestHandler: nil, // Is set below.
		},
		ResultExpected: SimpleTestResult{
			ResponseBodyString: "Some Error",
			ResponseStatusCode: http.StatusBadRequest,
		},
	}
	httpTest.Parameter.RequestHandler = func(w http.ResponseWriter, r *http.Request) {
		ReplyErrorWithCode( // <- This HTTP Handler Function is being tested.
			w,
			httpTest.ResultExpected.ResponseStatusCode,
			errors.New(httpTest.ResultExpected.ResponseBodyString),
		)
	}
	err = PerformSimpleHttpTest(&httpTest)
	test.MustBeNoError(err)
	test.MustBeEqual(httpTest.ResultReceived, httpTest.ResultExpected)
}

func Test_ReplyErrorInternal(t *testing.T) {

	var err error
	var httpTest SimpleTest
	var test *tester.Test

	test = tester.New(t)
	httpTest = SimpleTest{
		Parameter: SimpleTestParameter{
			RequestMethod:  TestMethod,
			RequestUrl:     TestUrl,
			RequestBody:    nil,
			RequestHandler: nil, // Is set below.
		},
		ResultExpected: SimpleTestResult{
			ResponseBodyString: "Some Internal Error",
			ResponseStatusCode: http.StatusInternalServerError,
		},
	}
	httpTest.Parameter.RequestHandler = func(w http.ResponseWriter, r *http.Request) {
		ReplyErrorInternal( // <- This HTTP Handler Function is being tested.
			w,
			errors.New(httpTest.ResultExpected.ResponseBodyString),
		)
	}
	err = PerformSimpleHttpTest(&httpTest)
	test.MustBeNoError(err)
	test.MustBeEqual(httpTest.ResultReceived, httpTest.ResultExpected)
}

func Test_ReplyJSON(t *testing.T) {

	type TestObjectClass struct {
		Age  int    `json:"age"`
		Name string `json:"name"`
	}

	var err error
	var httpTest SimpleTest
	var test *tester.Test

	test = tester.New(t)
	httpTest = SimpleTest{
		Parameter: SimpleTestParameter{
			RequestMethod:  TestMethod,
			RequestUrl:     TestUrl,
			RequestBody:    nil,
			RequestHandler: nil, // Is set below.
		},
		ResultExpected: SimpleTestResult{
			ResponseBodyString: `{"age":123,"name":"John"}`,
			ResponseStatusCode: http.StatusOK,
		},
	}
	testObject := TestObjectClass{
		Age:  123,
		Name: "John",
	}
	httpTest.Parameter.RequestHandler = func(w http.ResponseWriter, r *http.Request) {
		ReplyJSON( // <- This HTTP Handler Function is being tested.
			w,
			testObject,
		)
	}
	err = PerformSimpleHttpTest(&httpTest)
	test.MustBeNoError(err)
	test.MustBeEqual(httpTest.ResultReceived, httpTest.ResultExpected)
}

func Test_ReplyJSONFast(t *testing.T) {

	type TestObjectClass struct {
		Age  int    `json:"age"`
		Name string `json:"name"`
	}

	var err error
	var httpTest SimpleTest
	var test *tester.Test

	test = tester.New(t)
	httpTest = SimpleTest{
		Parameter: SimpleTestParameter{
			RequestMethod:  TestMethod,
			RequestUrl:     TestUrl,
			RequestBody:    nil,
			RequestHandler: nil, // Is set below.
		},
		ResultExpected: SimpleTestResult{
			ResponseBodyString: `{"age":123,"name":"John"}` + "\n", //!
			ResponseStatusCode: http.StatusOK,
		},
	}
	testObject := TestObjectClass{
		Age:  123,
		Name: "John",
	}
	httpTest.Parameter.RequestHandler = func(w http.ResponseWriter, r *http.Request) {
		ReplyJSONFast( // <- This HTTP Handler Function is being tested.
			w,
			testObject,
		)
	}
	err = PerformSimpleHttpTest(&httpTest)
	test.MustBeNoError(err)
	test.MustBeEqual(httpTest.ResultReceived, httpTest.ResultExpected)
}
