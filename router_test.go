package routing

import (
	"net/http"
	"testing"
)

var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

func Test_NewRouter_validRoutes(t *testing.T) {
	routes := Routes{}
	for _, verb := range []string{"GET", "PUT", "POST", "PATCH", "DELETE", "HEAD", "TRACE", "OPTIONS", "CONNECT"} {
		routes = append(routes, Route{verb, "/example1", handler})
	}

	router, err := NewRouter(routes)
	if router == nil {
		t.Errorf("expected instance of *mux.Router, got nil")
	}
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func Test_NewRouter_invalidHTTPVerbs(t *testing.T) {
	routes := Routes{
		{"NOPE", "/notagoodverb", handler},
	}

	var err error
	if _, err = NewRouter(routes); err == nil {
		t.Fatal("expected error, got nil")
	}

	expected := "routing: Route NOPE /notagoodverb does not have a valid HTTP verb"
	if err.Error() != expected {
		t.Fatalf("expected error: %s, got: %v", expected, err)
	}
}

func Test_NewRouter_duplicatePaths(t *testing.T) {
	routes := Routes{
		{"GET", "/example1", handler},
		{"GET", "/example1", handler},
	}

	var err error
	if _, err = NewRouter(routes); err == nil {
		t.Fatal("expected error, got nil")
	}

	expected := "routing: Route GET /example1 is defined more than once"
	if err.Error() != expected {
		t.Fatalf("expected error: %s, got: %v", expected, err)
	}
}
