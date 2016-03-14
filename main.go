package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Ibnul Cevzi Library Manager")
	var db = new(mongo)
	router := NewRouter(db)

	log.Fatal(http.ListenAndServe(":8080", router))
}
