package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Ibnul Cevzi Library Manager")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
