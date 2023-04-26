package main

import (
	"fmt"
	"io.analog.alex.mockserver/pkg/parser"
	"io.analog.alex.mockserver/pkg/server"
	"os"
	"strconv"
)

const (
	name   = "MockServer"
	author = "Miguel Alexandre"
	at     = "@2022"
)

var (
	port        = 8080
	defaultPath = "endpoints.json"
)

func main() {
	fmt.Printf("%s made with ❤️ by %s %s \n", name, author, at)

	// load context from ENV VARS if they are present
	loadContext()

	fmt.Printf("Start server at '%d' with endpoints file: (%s)\n", port, defaultPath)
	server.RunServer(
		fmt.Sprintf(":%d", port) /* http lib consumes port as string */, parser.LoadEndpointsFromContext(defaultPath),
	)
}

func loadContext() {
	// check for ENV VARS
	var contextPort = os.Getenv("SERVER_PORT")
	var contextPath = os.Getenv("ENDPOINTS_FILE_PATH")

	// set port from SERVER_PORT if env var is present
	if contextPort != "" {
		// handle conversion of string to int for SERVER_PORT field
		i, err := strconv.Atoi(contextPort)
		if err != nil {
			fmt.Printf("Could not convert %s to int, using default value %d\n", contextPort, port)
		} else {
			port = i
		}
	}

	// set path from ENDPOINTS_FILE_PATH if env var is present
	if contextPath != "" {
		defaultPath = contextPath
	}
}
