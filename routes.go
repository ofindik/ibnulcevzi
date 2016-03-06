package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"BookGetAll",
		"GET",
		"/ibnulcevzi/v1/books",
		BookGetAll,
	},
	Route{
		"BookGet",
		"GET",
		"/ibnulcevzi/v1/books/{bookId}",
		BookGet,
	},
	Route{
		"BookPost",
		"POST",
		"/ibnulcevzi/v1/books",
		BookPost,
	},
	Route{
		"BookPut",
		"PUT",
		"/ibnulcevzi/v1/books",
		BookPut,
	},
	Route{
		"BookDelete",
		"DELETE",
		"/ibnulcevzi/v1/books/{bookId}",
		BookDelete,
	},
}
