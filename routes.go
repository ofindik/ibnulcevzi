package main

import (
	"net/http"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.Handler
}

type Routes []Route

var db = new(mongo)

var routes = Routes{
	Route{
		"BookGetAll",
		"GET",
		"/ibnulcevzi/v1/books",
		BookGetAll(db),
	},
	Route{
		"BookGet",
		"GET",
		"/ibnulcevzi/v1/books/{bookId}",
		BookGet(db),
	},
	Route{
		"BookPost",
		"POST",
		"/ibnulcevzi/v1/books",
		BookPost(db),
	},
	Route{
		"BookPut",
		"PUT",
		"/ibnulcevzi/v1/books",
		BookPut(db),
	},
	Route{
		"BookDelete",
		"DELETE",
		"/ibnulcevzi/v1/books/{bookId}",
		BookDelete(db),
	},
}
