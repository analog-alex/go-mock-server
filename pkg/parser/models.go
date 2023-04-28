package parser

type Request struct {
	QueryParams map[string][]string `json:"query_params"`
	Headers     map[string][]string
	Body        string
}

type Response struct {
	StatusCode int `json:"status_code"`
	Body       string
}

type Endpoint struct {
	Uri      string
	Request  *Request
	Response Response
}

type Endpoints struct {
	Name        string
	Version     string
	Description string
	Authors     []string
	Endpoints   []Endpoint
}
