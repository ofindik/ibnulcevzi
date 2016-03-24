package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(db AppDatabase) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	var handler http.Handler
	handler = BookGetAll(db)
	handler = Logger(handler, "BookGetAll")
	router.
		Methods("GET").
		Path("/ibnulcevzi/v1/books").
		Name("BookGetAll").
		Handler(handler)

	handler = BookGet(db)
	handler = Logger(handler, "BookGet")
	router.
		Methods("GET").
		Path("/ibnulcevzi/v1/books/{bookId}").
		Name("BookGet").
		Handler(handler)

	handler = BookPost(db)
	handler = Logger(handler, "BookPost")
	router.
		Methods("POST").
		Path("/ibnulcevzi/v1/books").
		Name("BookPost").
		Handler(handler)

	handler = BookPut(db)
	handler = Logger(handler, "BookPut")
	router.
		Methods("PUT").
		Path("/ibnulcevzi/v1/books").
		Name("BookPut").
		Handler(handler)

	handler = BookDelete(db)
	handler = Logger(handler, "BookDelete")
	router.
		Methods("DELETE").
		Path("/ibnulcevzi/v1/books/{bookId}").
		Name("BookDelete").
		Handler(handler)

	return router
}
