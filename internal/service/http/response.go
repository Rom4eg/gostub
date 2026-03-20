package http

import "net/http"

type Response struct {
	Body   []byte
	Header http.Header
	Code   int
}
