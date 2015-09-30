package routing

import "net/http"

type Route struct {
	Method  string
	Path    string
	Handler http.Handler
}

func (rt Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rt.Handler.ServeHTTP(w, r)
}
