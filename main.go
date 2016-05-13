package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Ibnul Cevzi Library Manager")
	dbPtr := flag.String("database", "postgres", "Database to persist data")
	fmt.Println("Database flag:" + *dbPtr)
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	var db AppDatabase
	if *dbPtr == "postgres" {
		db = new(postgres)
	} else if *dbPtr == "mongo" {
		db = new(mongo)
	}

	router := NewRouter(db)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
