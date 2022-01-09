package parser

import (
    "mockServer/src/parser"
    "testing"
)

const (
    path1 string = "../../test/resources/endpoints_test.json"
)

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
}
