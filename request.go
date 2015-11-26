package main

const (
	GET = "GET"
	POST = "POST"
	DELETE = "DELETE"
	PUT = "PUT"
)

type CCARequest struct {
	Method string
	Endpoint string
	Body []byte
	Options map[string]string
}