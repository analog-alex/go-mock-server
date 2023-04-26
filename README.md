# GoMock Server

![Logo](img/go_mock.png)

A simple mock server configurable via JSON, built using GoLang. Go gophers!

## Run it locally

Checkout the project with **git clone**.
Assuming Go 1.19+ is installed on your machine, run the following commands:

```bash
export ENDPOINTS_FILE_PATH=json_examples/endpoints.json
go run cmd/main.go 
````

Change the `ENDPOINTS_FILE_PATH` environment variable to point to your endpoints file.

## How to create endpoints

A file name `endpoint.json` must be placed in the context root, or in the `ENDPOINTS_FILE_PATH` environment variable
with the following structure:

```json
{
  "endpoints": [
    {
      "uri": "/test",
      "response": {
        "status_code": 200,
        "body": "{ \"payload\": \"This is a response body!\" }"
      }
    }
  ]
}
```

As it stands you can match an endpoint, headers and request body and the server will return the configured response
if a match is found.

## Run tests

Simply run on project root:

```bash
go test ./test/...
```

by analog-alex (Miguel Alexandre) @2022
