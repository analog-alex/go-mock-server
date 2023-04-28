package parser

import (
	"github.com/stretchr/testify/assert"
	"io.analog.alex.mockserver/pkg/parser"
	"reflect"
	"testing"
)

const (
	path1 string = "../../test/resources/endpoints_test.json"
	path2 string = "../../test/resources/endpoints_request_test.json"
)

/*
Tests for the parser package
*/
func TestParserMetadata(t *testing.T) {
	a := assert.New(t)

	result := parser.LoadEndpointsFromContext(path1)

	// assert metadata
	a.Equal("1.0.0", result.Version)
	a.Equal("This is a test", result.Description)
	a.Equal([]string{"Miguel Alexandre"}, result.Authors)
}

func TestParser(t *testing.T) {
	a := assert.New(t)

	result := parser.LoadEndpointsFromContext(path1)
	a.Equal(1, len(result.Endpoints))

	// get the first (and only) route
	testSubject := result.Endpoints[0]
	a.Equal("/test", testSubject.Uri)

	// assert on response
	response := testSubject.Response
	a.Equal(202, response.StatusCode)
	a.Equal("{ \"payload\": \"This is a test!\" }", response.Body)

	// assert on request
	request := testSubject.Request
	a.Nilf(request, "Request structure should not be present")
}

func TestParserRequest(t *testing.T) {
	a := assert.New(t)

	result := parser.LoadEndpointsFromContext(path2)
	a.Equal(1, len(result.Endpoints))

	// get the first (and only) route with a request
	testSubject := result.Endpoints[0].Request
	a.NotNilf(testSubject, "Request structure should be present")

	// assert on query params
	paramsExpected := make(map[string][]string)
	paramsExpected["name"] = []string{"Miguel"}

	a.True(reflect.DeepEqual(paramsExpected, testSubject.QueryParams), "Query params did not match")

	// assert on headers
	headersExpected := make(map[string][]string)
	headersExpected["Authorization"] = []string{"token"}

	a.True(reflect.DeepEqual(headersExpected, testSubject.Headers), "Headers did not match")

	// assert on body
	a.Equal("{ \"payload\": \"This is a body!\" }", testSubject.Body)
}
