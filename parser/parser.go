package parser

import (
    "encoding/json"
    "fmt"
    "os"
)

const (
    path = "endpoints.json"
)

func parseJsonFile() string {
    data, err := os.ReadFile("endpoints.json")
    if err != nil {
        panic(fmt.Sprintf("No file %s found to load", path))
    }

    // return raw JSON as a string
    return string(data)
}

func LoadEndpointsFromContext() []Endpoint {
    content := parseJsonFile()
    var endpoints Endpoints

    err := json.Unmarshal([]byte(content), &endpoints)
    if err != nil {
        panic(fmt.Sprintf("Could not parse JSON in %s to valid endpoints. Check documentation.", path))
    }

    return endpoints.Endpoints
}
