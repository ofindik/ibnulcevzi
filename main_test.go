package main

import (
	"bytes"
	"gopkg.in/gorilla/mux.v0"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var bookList map[string]Book

type MockDb struct {
}

func (m MockDb) Readall(url, dbName, cName string) (books Books) {
	return Books{Book{Name: "TestBook1", Author: "TestAuthor1"}}
}

func (m MockDb) Read(url, dbName, cName, bookId string) Book {
	book, ok := bookList[bookId]
	if !ok {
		book = Book{Name: "TestBook2", Author: "TestAuthor2"}
	}
	return book
}

func (m MockDb) Write(url, dbName, cName string, b Book) {
	bookList[b.Name] = b
}

func (m MockDb) Update(url, dbName, cName string, b Book) {
	bookList[b.Name] = b
}

func (m MockDb) Delete(url, dbName, cName, bookId string) {
}

func init() {
	bookList = make(map[string]Book)
}

func TestGetAll(t *testing.T) {
	mockDb := MockDb{}

	req, _ := http.NewRequest("GET", "/ibnulcevzi/v1/books", nil)
	w := httptest.NewRecorder()

	NewRouter(mockDb).ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Get all books didn't return StatusOK: %v", w.Code)
	}
	books := w.Body.String()
	if !strings.Contains(books, "TestBook1") || !strings.Contains(books, "TestAuthor1") {
		t.Errorf("Get all books returned unexpected result: %v", books)
	}
}

func TestGet(t *testing.T) {
	mockDb := MockDb{}

	req, _ := http.NewRequest("GET", "/ibnulcevzi/v1/books/TestBook2", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	NewRouter(mockDb).ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Get book didn't return StatusOK: %v", w.Code)
	}
	books := w.Body.String()
	if !strings.Contains(books, "TestBook2") || !strings.Contains(books, "TestAuthor2") {
		t.Errorf("Get book returned unexpected result: %v", books)
	}
}

func TestPost(t *testing.T) {
	mockDb := MockDb{}

	var jsonStr = []byte(`{"ID":"56da1253deff5fe5697fac9a","name":"TestBook3","author":"TestAuthor3","status":"","borrowed":"0001-01-01T00:00:00Z"}`)

	req, _ := http.NewRequest("POST", "/ibnulcevzi/v1/books", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	NewRouter(mockDb).ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Post book didn't return http.StatusCreated: %v", w.Code)
	}
	books := w.Body.String()
	if !strings.Contains(books, "TestBook3") || !strings.Contains(books, "TestAuthor3") {
		t.Errorf("Post book returned unexpected result: %v", books)
	}
}

func TestPut(t *testing.T) {
	mockDb := MockDb{}

	var jsonStr = []byte(`{"ID":"56da1253deff5fe5697fac9a","name":"TestBook4","author":"UnknownAuthor","status":"","borrowed":"0001-01-01T00:00:00Z"}`)

	req, _ := http.NewRequest("PUT", "/ibnulcevzi/v1/books", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	NewRouter(mockDb).ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Put book didn't return http.StatusOK: %v", w.Code)
	}
	books := w.Body.String()
	if !strings.Contains(books, "TestBook4") || !strings.Contains(books, "UnknownAuthor") {
		t.Errorf("Put book returned unexpected result: %v", books)
	}
}

func TestDelete(t *testing.T) {
	mockDb := MockDb{}

	req, _ := http.NewRequest("DELETE", "/ibnulcevzi/v1/books/TestBook6", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	NewRouter(mockDb).ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Delete book didn't return http.StatusOK: %v", w.Code)
	}
}

func TestPostAndGet(t *testing.T) {
	mockDb := MockDb{}

	var jsonStr = []byte(`{"ID":"56da1253deff5fe5697fac9a","name":"TestBook4","author":"TestAuthor4","status":"","borrowed":"0001-01-01T00:00:00Z"}`)

	req, _ := http.NewRequest("POST", "/ibnulcevzi/v1/books", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	var testRouter *mux.Router = NewRouter(mockDb)
	testRouter.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Post book didn't return http.StatusCreated: %v", w.Code)
	}
	books := w.Body.String()
	if !strings.Contains(books, "TestBook4") || !strings.Contains(books, "TestAuthor4") {
		t.Errorf("Post book returned unexpected result: %v", books)
	}

	req, _ = http.NewRequest("GET", "/ibnulcevzi/v1/books/TestBook4", nil)
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	testRouter.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Get book didn't return StatusOK: %v", w.Code)
	}
	books = w.Body.String()
	if !strings.Contains(books, "TestBook4") || !strings.Contains(books, "TestAuthor4") {
		t.Errorf("Get book returned unexpected result: %v", books)
	}
}
