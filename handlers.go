package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func BookIndex(w http.ResponseWriter, r *http.Request) {
	books := Books{
		Book{Name: "Book1", Author: "Author1"},
		Book{Name: "Book2", Author: "Author2"},
		Book{Name: "Book3", Author: "Author3"},
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(books); err != nil {
		panic(err)
	}
}

func BookShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	fmt.Fprintf(w, "Todo show:", bookId)
}
