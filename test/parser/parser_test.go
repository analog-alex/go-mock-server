package parser

import (
	"io.analog.alex.mockserver/pkg/parser"
	"reflect"
	"testing"
)

const (
	path1 string = "../../test/resources/endpoints_test.json"
	path2 string = "../../test/resources/endpoints_request_test.json"
)

/*
Test simple JSON endpoint with no request content
*/
func TestParser(t *testing.T) {
	result := parser.LoadEndpointsFromContext(path1)

	if len(result) != 1 {
		t.Fatalf("Parsed file did not retrieve one endpoint")
	}

	testSubject := result[0]

	if testSubject.Uri != "/test" {
		t.Fatalf("Endpoint uri %s was not '/test'", testSubject.Uri)
	}

	response := testSubject.Response

	if response.StatusCode != 202 {
		t.Fatalf("Status code %d was not '202'", response.StatusCode)
	}

	if response.Body != "{ \"payload\": \"This is a test!\" }" {
		t.Fatalf("Status code %s was not '\"{ \"payload\": \"This is a test!\" }\"'", response.Body)
	}

	if testSubject.Request != nil {
		t.Fatalf("Request structure should not be present")
	}
}

/*
Test JSON endpoint with defined request content
*/
func TestParserWithRequest(t *testing.T) {
	result := parser.LoadEndpointsFromContext(path2)

	if len(result) != 1 {
		t.Fatalf("Parsed file did not retrieve one endpoint")
	}

	testSubject := result[0].Request

	if testSubject == nil {
		t.Fatalf("Request structure should be present")
	}

	// check query params
	params := make(map[string][]string)
	params["name"] = []string{"Miguel"}

	if !reflect.DeepEqual(testSubject.QueryParams, params) {
		t.Fatalf("QueryParams map did not match %s", params)
	}

	// check headers
	headers := make(map[string][]string)
	headers["Authorization"] = []string{"token"}

	if !reflect.DeepEqual(testSubject.Headers, headers) {
		t.Fatalf("Header map did not match %s", headers)
	}

	// check body
	if testSubject.Body != "{ \"payload\": \"This is a body!\" }" {
		t.Fatalf("Status code %s was not '{ \"payload\": \"This is a body!\" }'", testSubject.Body)
	}
}
