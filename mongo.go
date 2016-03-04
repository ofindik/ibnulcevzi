package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
)

var session *mgo.Session
var c *mgo.Collection

func connect() {
	log.Printf(
		"Will connect to local mongodb",
	)
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	log.Printf(
		"Connected to local mongodb",
	)
	session.SetMode(mgo.Monotonic, true)

	c = session.DB("ibnulcevzi").C("books")
	log.Printf(
		"Session retireved",
	)
}

func close() {
	session.Close()
}

func write(b Book) {
	oldBook := read(b.Name)
	if "" != oldBook.Name {
		log.Printf("Book already exist:", oldBook.Name)
	} else {
		err := c.Insert(b)
		if err != nil {
			panic(err)
		}
	}
}

func read(name string) Book {
	result := Book{}
	err := c.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		log.Printf("Book not found:", name)
	}
	return result
}
