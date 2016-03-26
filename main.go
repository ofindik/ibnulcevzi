package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Ibnul Cevzi Library Manager")
	portPtr := flag.String("port", "8080", "HTTP port to listen")
	var db = new(mongo)
	router := NewRouter(db)
	log.Fatal(http.ListenAndServe(":"+*portPtr, router))
}
