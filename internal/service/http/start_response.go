package http

import "net/http"

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

	_, err := w.Write(r.Body)
	if err != nil {
		s.l.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
}
