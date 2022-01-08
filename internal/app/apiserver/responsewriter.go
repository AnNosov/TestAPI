package apiserver

import "net/http"

type responseWriter struct {
	http.ResponseWriter
	code int
}

func (w *responseWriter) WriteHeader(StatusCode int) {
	w.code = StatusCode
	w.ResponseWriter.WriteHeader(StatusCode)
}
