package routing

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_ServeHTTP_callsHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := &http.Request{}
	handler := &testHandler{}

	route := Route{"", "", handler}
	route.ServeHTTP(w, r)

	if handler.w != w {
		t.Fatalf("ServeHTTP expected ResponseWriter %v, got %v", w, handler.w)
	}

	if handler.r != r {
		t.Fatalf("ServeHTTP expected Request %v, got %v", r, handler.r)
	}
}

type testHandler struct {
	w http.ResponseWriter
	r *http.Request
}

func (h *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.w = w
	h.r = r
}
