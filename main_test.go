package main

import (
	"bytes"
	"github.com/gorilla/context"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockDb struct{}

func (m MockDb) Readall(url, dbName, cName string) (books Books) {
	log.Printf("MockDb Readall")
	return Books{Book{Name: "TestBook1", Author: "TestAuthor1"}}
}

func (m MockDb) Read(url, dbName, cName, bookId string) Book {
	log.Printf("MockDb Read:", bookId)
	return Book{Name: "TestBook2", Author: "TestAuthor2"}
}

func (m MockDb) Write(url, dbName, cName string, b Book) {
	log.Printf("MockDb Write:", b)
}

func (m MockDb) Update(url, dbName, cName string, b Book) {
}

func (m MockDb) Delete(url, dbName, cName, bookId string) {
}

func TestGetAll(t *testing.T) {
	mockDb := MockDb{}
	testHandler := BookGetAll(mockDb)

	req, _ := http.NewRequest("GET", "/ibnulcevzi/v1/books", nil)
	w := httptest.NewRecorder()

	testHandler.ServeHTTP(w, req)
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
	testHandler := BookGet(mockDb)

	req, _ := http.NewRequest("GET", "/ibnulcevzi/v1/books/TestBook2", nil)
	req.Header.Set("Content-Type", "application/json")
	context.Set(req, "bookId", "TestBook222")
	w := httptest.NewRecorder()

	testHandler.ServeHTTP(w, req)
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
	testHandler := BookPost(mockDb)

	var jsonStr = []byte(`{"ID":"56da1253deff5fe5697fac9a","name":"TestBook3","author":"TestAuthor3","status":"","borrowed":"0001-01-01T00:00:00Z"}`)

	req, _ := http.NewRequest("POST", "/ibnulcevzi/v1/books", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	testHandler.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Post book didn't return http.StatusCreated: %v", w.Code)
	}
	books := w.Body.String()
	if !strings.Contains(books, "TestBook3") || !strings.Contains(books, "TestAuthor3") {
		t.Errorf("Post book returned unexpected result: %v", books)
	}
}
