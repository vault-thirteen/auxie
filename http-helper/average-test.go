package httphelper

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/vault-thirteen/auxie/errors"
)

// An average test can emulate an HTTP request with HTTP method, URL, HTTP
// headers and body.
//
// It executes the HTTP handler specified in the 'RequestHandler' field,
// provides the results of this execution in a field named 'ResultReceived'.
// A user may then compare the received results with the expected ones.

// AverageTest is an average HTTP test.
type AverageTest struct {
	Parameter      AverageTestParameter
	ResultExpected AverageTestResult
	ResultReceived AverageTestResult
}

// AverageTestParameter is a parameter of an average HTTP test.
type AverageTestParameter struct {
	RequestMethod  string
	RequestUrl     string
	RequestHeaders http.Header
	RequestBody    io.Reader
	RequestHandler http.HandlerFunc
}

// AverageTestResult is a result of an average HTTP test.
type AverageTestResult struct {
	ResponseStatusCode int
	ResponseHeaders    http.Header
	ResponseBody       []byte
}

// PerformAverageHttpTest function performs the simulation of an average HTTP
// test handler. It writes the received results into the 'ResultReceived' field
// of a test object.
func PerformAverageHttpTest(test *AverageTest) (err error) {

	// Prepare data.
	request := httptest.NewRequest(
		test.Parameter.RequestMethod,
		test.Parameter.RequestUrl,
		test.Parameter.RequestBody,
	)
	request.Header = test.Parameter.RequestHeaders

	responseRecorder := httptest.NewRecorder()

	// Make a simulated request to an HTTP handler.
	test.Parameter.RequestHandler(responseRecorder, request)

	// Get the response.
	response := responseRecorder.Result()

	var responseBody []byte
	responseBody, err = io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	defer func() {
		derr := response.Body.Close()
		if derr != nil {
			err = errors.Combine(err, derr)
		}
	}()

	// Set the result.
	test.ResultReceived = AverageTestResult{
		ResponseStatusCode: response.StatusCode,
		ResponseHeaders:    response.Header,
		ResponseBody:       responseBody,
	}

	return nil
}
