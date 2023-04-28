package parser

import (
	"encoding/json"
	"fmt"
	"os"
)

func getJsonByFileName(fileName string) string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(fmt.Sprintf("No file %s found to load", fileName))
	}

	// return raw JSON as a string
	return string(data)
}

func LoadEndpointsFromContext(fileName string) Root {
	var endpoints Root
	content := getJsonByFileName(fileName)

	if err := json.Unmarshal([]byte(content), &endpoints); err != nil {
		// can't do anything with a bad file so we panic
		panic(fmt.Sprintf("Could not parse JSON in %s to valid endpoints. Check documentation.", fileName))
	}

	return endpoints
}
