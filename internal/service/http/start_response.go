package http

import (
	"fmt"
	"net/http"
)

func (s *Service) StartResponse(r *Response, w http.ResponseWriter) {
	s.l.Debug("Enter StartResponse")
	defer s.l.Debug("Exit StartResponse")

	for k, v := range r.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}

	code := r.Code
	if code < 100 {
		code = http.StatusNotImplemented
		if len(r.Body) > 0 {
			code = http.StatusOK
		}
	}
	w.WriteHeader(code)

	n, err := w.Write(r.Body)
	if err != nil {
		s.l.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)

		_, err = w.Write([]byte(err.Error()))
		s.l.Error(err.Error())
		return
	}

	if n < len(r.Body) {
		s.l.Error(fmt.Sprintf("writen %d bytes, while recieved %d bytes", n, len(r.Body)))
	}
}
