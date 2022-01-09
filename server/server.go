package server

import (
    "fmt"
    "mockServer/parser"
    "net/http"
)

func RunServer(port string, endpoints []parser.Endpoint) bool {
    for _, endpoint := range endpoints {
        // build and bind the handler for each endpoint
        http.HandleFunc(endpoint.Uri, createHandlerForEndpoint(endpoint))
    }

    err := http.ListenAndServe(fmt.Sprintf("%s", port), nil /* nothing goes here */)
    if err != nil {
        _ = fmt.Errorf("server failed to boot")
        return false
    }

    return true
}

func createHandlerForEndpoint(endpoint parser.Endpoint) func(http.ResponseWriter, *http.Request) {
    return func(writer http.ResponseWriter, request *http.Request) {
        // set Content-Type header
        writer.Header().Set("Content-Type", "application/json")

        // set the response status code BEFORE writing anything into the writer
        writer.WriteHeader(endpoint.Response.StatusCode)

        // output the body
        _, err := fmt.Fprintf(writer, endpoint.Response.Body)
        if err != nil {
            panic(fmt.Sprintf("Error at endpoint %s", endpoint.Uri))
        }
    }
}