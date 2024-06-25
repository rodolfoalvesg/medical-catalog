package responses

import "net/http"

type Response struct {
	Status  int
	header  http.Header
	Payload interface{}
	Error   error
}

func (r Response) Header() http.Header {
	return r.header
}
