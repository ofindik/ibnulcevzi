package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	DB_HOST     = "ec2-176-34-127-73.eu-west-1.compute.amazonaws.com"
	DB_USER     = "lckergiimctodl"
	DB_PASSWORD = "hxs38v05nnnDZZMcwkaQjsBBel"
	DB_NAME     = "d696d17afrfe9h"
)

type postgres struct{}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func connectPostGres() (db *sql.DB, err error) {
	//dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
	//DB_HOST, DB_USER, DB_PASSWORD, DB_NAME)
	connectionURL := "postgres://xvbklnzozgakdk:QNHaBxdLo8CUrrTL8g--KuMXkx@ec2-176-34-127-73.eu-west-1.compute.amazonaws.com:5432/dtrgpf9tvql9e"
	db, err = sql.Open("postgres", connectionURL)
	checkErr(err)
	log.Printf("Connected to: ", DB_NAME)
	return
}

func closePostgres(db *sql.DB) {
	db.Close()
	log.Printf("Database closed")
}

func (p *postgres) Readall(url, dbName, cName string) (books Books) {
	db, err := connectPostGres()
	defer closePostgres(db)
	rows, err := db.Query("SELECT * FROM book")
	checkErr(err)
	result := Books{}
	fmt.Println("id | name | author | status")
	for rows.Next() {
		var id int
		var name string
		var author string
		var status string
		err = rows.Scan(&id, &name, &author, &status)
		checkErr(err)
		fmt.Printf("%3v | %8v | %6v | %6v \n", id, name, author, status)
		book := Book{Name: name, Author: author, Status: status}
		result = append(result, book)
	}
	return result
}

func (p *postgres) Read(url, dbName, cName, bookId string) Book {
	result := Book{}

	return result
}
func (p *postgres) Write(url, dbName, cName string, b Book) {
}
func (p *postgres) Update(url, dbName, cName string, b Book) {
}
func (p *postgres) Delete(url, dbName, cName, bookId string) {
}
