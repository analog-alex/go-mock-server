package parser

type Response struct {
    StatusCode int `json:"status_code"`
    Body       string
}

type Endpoint struct {
    Uri      string
    Response Response
}

type Endpoints struct {
    Endpoints []Endpoint
}
