package main

import (
    "fmt"
    "mockServer/parser"
    "mockServer/server"
)

var (
    name   = "MockServer"
    port   = ":8080"
    author = "Miguel Alexandre"
    at     = "@2022"
)

func main() {
    fmt.Printf("%s made with ❤️ by %s %s \n", name, author, at)
    fmt.Printf("Start server at %s\n", port)
    server.RunServer(port, parser.LoadEndpointsFromContext())
}
