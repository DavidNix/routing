package routing

import (
	"fmt"

	"github.com/gorilla/mux"
)

var (
	verbs = map[string]bool{
		"GET":     true,
		"PUT":     true,
		"POST":    true,
		"PATCH":   true,
		"DELETE":  true,
		"HEAD":    true,
		"TRACE":   true,
		"OPTIONS": true,
		"CONNECT": true,
	}
)

type Routes []Route

func NewRouter(routes Routes) (*mux.Router, error) {
	router := mux.NewRouter()

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Path).
			Handler(route)
	}

	return router, validate(routes)
}

func validate(routes Routes) error {
	var err error
	counter := make(map[string]int)

	for _, route := range routes {
		if err = validateHTTPVerb(route); err != nil {
			return err
		}
		signature := fmt.Sprint(route.Method, route.Path)
		counter[signature]++
		if counter[signature] > 1 {
			return fmt.Errorf("routing: Route %s %s is defined more than once", route.Method, route.Path)
		}
	}
	return nil
}

func validateHTTPVerb(route Route) error {
	if !verbs[route.Method] {
		return fmt.Errorf("routing: Route %s %s does not have a valid HTTP verb", route.Method, route.Path)
	}
	return nil
}
