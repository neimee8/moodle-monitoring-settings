package wrappers

import (
	"bytes"
	"net/http"
)

type ResponseWriter struct {
	http.ResponseWriter
	Status      int
	Body        bytes.Buffer
	wroteHeader bool
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		Status:         200,
	}
}

func (w *ResponseWriter) WriteHeader(statusCode int) {
	if w.wroteHeader {
		return
	}

	w.Status = statusCode
	w.wroteHeader = true
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *ResponseWriter) Write(b []byte) (int, error) {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}

	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}
