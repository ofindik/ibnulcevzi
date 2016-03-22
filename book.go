package main

import (
	"gopkg.in/mgo.v2/bson"
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
