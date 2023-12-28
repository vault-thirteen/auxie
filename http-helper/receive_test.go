package httphelper

import (
	"net/http"
	"strings"
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_ReceiveJSON(t *testing.T) {
	const TestedHttpBody = `{"age":12345,"name":"Decode me"}`

	type TestObjectClass struct {
		Age  int    `json:"age"`
		Name string `json:"name"`
	}

	var err error
	var httpTest SimpleTest
	var test = tester.New(t)

	httpTest = SimpleTest{
		Parameter: SimpleTestParameter{
			RequestMethod:  TestMethod,
			RequestUrl:     TestUrl,
			RequestBody:    nil, // Is set below.
			RequestHandler: nil, // Is set below.
		},
	}
	objectExpected := TestObjectClass{
		Age:  12345,
		Name: "Decode me",
	}

	// Test #1. Negative test: Not a pointer.
	// This HTTP handler receives an object and checks it.
	httpTest.Parameter.RequestBody = strings.NewReader(TestedHttpBody)
	httpTest.Parameter.RequestHandler = func(rw http.ResponseWriter, req *http.Request) {
		var handlerObject TestObjectClass
		herr := ReceiveJSON(req, handlerObject)
		test.MustBeAnError(herr)
	}
	err = PerformSimpleHttpTest(&httpTest)
	test.MustBeNoError(err)

	// Test #2. Positive test.
	// This HTTP handler receives an object and checks it.
	httpTest.Parameter.RequestBody = strings.NewReader(TestedHttpBody)
	httpTest.Parameter.RequestHandler = func(rw http.ResponseWriter, req *http.Request) {
		var handlerObject TestObjectClass
		herr := ReceiveJSON(req, &handlerObject)
		test.MustBeNoError(herr)
		test.MustBeEqual(handlerObject, objectExpected)
	}
	err = PerformSimpleHttpTest(&httpTest)
	test.MustBeNoError(err)
}
