package server

import (
	"fmt"
	"io.analog.alex.mockserver/pkg/parser"
	"io.analog.alex.mockserver/pkg/server/matchers"
	"log"
	"net/http"
)

func RunServer(port string, endpoints []parser.Endpoint) bool {
	for _, endpoint := range endpoints {
		// build and bind the handler for each endpoint
		http.HandleFunc(endpoint.Uri, createHandlerForEndpoint(endpoint))
	}

	err := http.ListenAndServe(fmt.Sprintf("%s", port), nil /* nothing goes here */)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func createHandlerForEndpoint(endpoint parser.Endpoint) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		if endpoint.Request != nil {
			// compare query params
			if !matchers.AreMapsEqual(request.URL.Query(), endpoint.Request.QueryParams) {
				err := fmt.Errorf("query params did not match. Expected %T got %T\n", endpoint.Request.QueryParams, request.URL.Query())
				log.Println(err)
				writer.WriteHeader(404)
				return
			}

			// compare headers
			if !matchers.IsMapASubset(request.Header, endpoint.Request.Headers) {
				err := fmt.Errorf("headers did not match. Expected %T got %T\n", endpoint.Request.Headers, request.Header)
				log.Println(err)
				writer.WriteHeader(404)
				return
			}
		}

		// set Content-Type header
		writer.Header().Set("Content-Type", "application/json")

		// set the response status code BEFORE writing anything into the writer
		writer.WriteHeader(endpoint.Response.StatusCode)

		// output the body
		_, err := fmt.Fprintf(writer, endpoint.Response.Body)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(500)
			return
		}

		log.Printf("Handled request %s\n", request.URL)
	}
}
