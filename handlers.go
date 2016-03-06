package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func BookGetAll(w http.ResponseWriter, r *http.Request) {
	connect()
	defer close()
	books := readall()

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(books); err != nil {
		panic(err)
	}
}

func retrieveBookId(r *http.Request) string {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	log.Printf("BookGet bookId", bookId)

	return bookId
}

func BookGet(w http.ResponseWriter, r *http.Request) {
	bookId := retrieveBookId(r)

	connect()
	defer close()
	book := read(bookId)

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(book); err != nil {
		panic(err)
	}
}

func readBookContent(w http.ResponseWriter, r *http.Request) Book {
	var book Book
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &book); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	return book
}

func BookPost(w http.ResponseWriter, r *http.Request) {
	book := readBookContent(w, r)

	connect()
	defer close()
	write(book)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(book); err != nil {
		panic(err)
	}
}

func BookPut(w http.ResponseWriter, r *http.Request) {
	book := readBookContent(w, r)

	connect()
	defer close()
	update(book)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(book); err != nil {
		panic(err)
	}
}

func BookDelete(w http.ResponseWriter, r *http.Request) {
	bookId := retrieveBookId(r)

	connect()
	defer close()
	delete(bookId)

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
