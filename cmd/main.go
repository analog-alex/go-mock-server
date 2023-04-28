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
	defaultPort     = 8080
	defaultHostname = "localhost"
	defaultPath     = "endpoints.json"
)

func main() {
	fmt.Printf("%s made with ❤️ by %s %s \n", name, author, at)

	// load context from ENV VARS if they are present -- this will set default values to config values
	loadEnv()

	fmt.Printf("Start server at '%d' with endpoints file: (%s)\n", defaultPort, defaultPath)
	server.RunServer(
		fmt.Sprintf("%s:%d", defaultHostname, defaultPort), parser.LoadEndpointsFromContext(defaultPath),
	)
}

func loadEnv() {
	// check for ENV VARS
	var port = os.Getenv("SERVER_PORT")
	var hostname = os.Getenv("SERVER_HOSTNAME")
	var contextPath = os.Getenv("ENDPOINTS_FILE_PATH")

	// set port from SERVER_PORT if env var is present
	if port != "" {
		// handle conversion of string to int for SERVER_PORT field
		i, err := strconv.Atoi(port)
		if err != nil {
			fmt.Printf("Could not convert %s to int, using default value %d\n", port, defaultPort)
		} else {
			defaultPort = i
		}
	}

	// set hostname from SERVER_HOSTNAME if env var is present
	if hostname != "" {
		defaultHostname = hostname
	}

	// set path from ENDPOINTS_FILE_PATH if env var is present
	if contextPath != "" {
		defaultPath = contextPath
	}
}
