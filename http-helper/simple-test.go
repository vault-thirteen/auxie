package httphelper

import (
	"io"
	"net/http"
	"net/http/httptest"
)

// SimpleTest is a simple HTTP test.
type SimpleTest struct {
	Parameter      SimpleTestParameter
	ResultExpected SimpleTestResult
	ResultReceived SimpleTestResult
}

// SimpleTestParameter is a parameter of a simple HTTP test.
type SimpleTestParameter struct {
	RequestMethod  string
	RequestUrl     string
	RequestBody    io.Reader
	RequestHandler http.HandlerFunc
}

// SimpleTestResult is a result of a simple HTTP test.
type SimpleTestResult struct {
	ResponseStatusCode int
	ResponseBodyString string
}

// PerformSimpleHttpTest function performs the simulation of a simple HTTP test
// handler. It writes the received results into the 'ResultReceived' field of a
// test object.
func PerformSimpleHttpTest(test *SimpleTest) (err error) {
	var request *http.Request
	var response *http.Response
	var responseBody []byte
	var responseRecorder *httptest.ResponseRecorder

	// Prepare Data.
	request = httptest.NewRequest(
		test.Parameter.RequestMethod,
		test.Parameter.RequestUrl,
		test.Parameter.RequestBody,
	)
	responseRecorder = httptest.NewRecorder()

	// Make a simulated Request to an HTTP Handler.
	test.Parameter.RequestHandler(responseRecorder, request)

	// Get the Response.
	response = responseRecorder.Result()
	responseBody, err = io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	err = response.Body.Close()
	if err != nil {
		return err
	}

	// Set the Result.
	test.ResultReceived = SimpleTestResult{
		ResponseBodyString: string(responseBody),
		ResponseStatusCode: response.StatusCode,
	}
	return nil
}
