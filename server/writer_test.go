package server

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestWriteJson(t *testing.T) {
	w := NewWriter()
	err := WriteJson(w, http.StatusOK, map[string]string{"hello": "world"})
	require.Nil(t, err)
}

func TestWriteJson_error(t *testing.T) {
	w := NewWriter()
	err := WriteJson(w, http.StatusOK, func() string{ return "this should error"})
	require.NotNil(t, err)
}

type testWriter struct {
	headers http.Header
	statusCode int
	body []byte
}

func NewWriter() *testWriter {
	return &testWriter{headers: http.Header{}}
}

func (w *testWriter) Header() http.Header {
	return w.headers
}

func (w *testWriter) Write(body []byte) (int, error){
	w.body = body
	return 0, nil
}

func (w *testWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}
