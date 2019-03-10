package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes slice
type Routes []Route

// NewRouter factory
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	router.NotFoundHandler = Logger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Nothing here: 404")
	}), "404")

	return router
}

var routes = Routes{
	Route{
		Name:        "Index",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: Index,
	},
	// Route{
	// 	Name:        "ListTodos",
	// 	Method:      "GET",
	// 	Pattern:     "/todos",
	// 	HandlerFunc: ListTodos,
	// },
	// Route{
	// 	Name:        "GetTodo",
	// 	Method:      "GET",
	// 	Pattern:     "/todos/{todoId}",
	// 	HandlerFunc: GetTodo,
	// },
}
