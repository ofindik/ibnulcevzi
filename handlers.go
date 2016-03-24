package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type AppDatabase interface {
	Readall(url, dbName, cName string) (books Books)
	Read(url, dbName, cName, bookId string) (book Book)
	Write(url, dbName, cName string, book Book)
	Update(url, dbName, cName string, book Book)
	Delete(url, dbName, cName, bookId string)
}

func BookGetAll(db AppDatabase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		books := db.Readall("mongodb://localhost:27017", "ibnulcevzi", "books")
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(books); err != nil {
			log.Printf("Error:", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
}

func retrieveBookId(r *http.Request) string {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	return bookId
}

func BookGet(db AppDatabase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bookId := retrieveBookId(r)
		book := db.Read("mongodb://localhost:27017", "ibnulcevzi", "books", bookId)
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(book); err != nil {
			log.Printf("Error:", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

	})
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
			log.Printf("Error:", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	return book
}

func BookPost(db AppDatabase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		book := readBookContent(w, r)
		db.Write("mongodb://localhost:27017", "ibnulcevzi", "books", book)
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(book); err != nil {
			log.Printf("Error:", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

	})
}

func BookPut(db AppDatabase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		book := readBookContent(w, r)
		db.Update("mongodb://localhost:27017", "ibnulcevzi", "books", book)
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(book); err != nil {
			log.Printf("Error:", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

	})
}

func BookDelete(db AppDatabase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bookId := retrieveBookId(r)
		db.Delete("mongodb://localhost:27017", "ibnulcevzi", "books", bookId)
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	})
}
