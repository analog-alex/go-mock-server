package main

import (
	"fmt"
	"mockServer/src/parser"
	"mockServer/src/server"
)

const (
	name   = "MockServer"
	author = "Miguel Alexandre"
	at     = "@2022"
)

var (
	port        = ":8080"
	defaultPath = "endpoints.json"
)

func main() {
	fmt.Printf("%s made with ❤️ by %s %s \n", name, author, at)
	fmt.Printf("Start server at %s\n", port)
	server.RunServer(port, parser.LoadEndpointsFromContext(defaultPath))
}
