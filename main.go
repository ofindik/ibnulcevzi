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
	dbPtr := flag.String("database", "postgres", "Daatabase to persist data")
	fmt.Println("Port flag:" + *portPtr)
	fmt.Println("Database flag:" + *dbPtr)
	var db AppDatabase
	if *dbPtr == "postgres" {
		db = new(postgres)
	} else if *dbPtr == "mongo" {
		db = new(mongo)
	}

	router := NewRouter(db)
	log.Fatal(http.ListenAndServe(":"+*portPtr, router))
}
