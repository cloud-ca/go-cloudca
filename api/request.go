package api

const (
	GET = "GET"
	POST = "POST"
	DELETE = "DELETE"
	PUT = "PUT"
)

type CcaRequest struct {
	Method string
	Endpoint string
	Body []byte
	Options map[string]string
}