package httphelper

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/vault-thirteen/auxie/errors"
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

	// Prepare Data.
	request := httptest.NewRequest(
		test.Parameter.RequestMethod,
		test.Parameter.RequestUrl,
		test.Parameter.RequestBody,
	)

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
	test.ResultReceived = SimpleTestResult{
		ResponseBodyString: string(responseBody),
		ResponseStatusCode: response.StatusCode,
	}

	return nil
}
