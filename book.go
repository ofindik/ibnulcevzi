package main

import "time"

type Book struct {
	Name     string    `json:"name"`
	Author   string    `json:"author"`
	Status   string    `json:"status"`
	Borrowed time.Time `json:"borrowed"`
}

type Books []Book
