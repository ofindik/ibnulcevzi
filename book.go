package main

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Book struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string        `json:"name"`
	Author   string        `json:"author"`
	Status   string        `json:"status"`
	Borrowed time.Time     `json:"borrowed"`
}

type Books []Book
