package server

import (
	"io.analog.alex.mockserver/pkg/parser"
	"testing"
)

const port = ":8080"

var endpoints = []parser.Endpoint{
	{
		Uri:     "/resource",
		Request: nil,
		Response: parser.Response{
			StatusCode: 200,
			Body:       "a JSON string",
		},
	},
	{
		Uri: "/other",
		Request: &parser.Request{
			QueryParams: map[string][]string{"key": {"value"}},
			Headers:     map[string][]string{"auth": {"token"}},
			Body:        "",
		},
		Response: parser.Response{
			StatusCode: 201,
			Body:       "",
		},
	},
}

func TestServerCorrectness(t *testing.T) {
	//server.RunServer(port, endpoints)

	// TODO make requests
}
