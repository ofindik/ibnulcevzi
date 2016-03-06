package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
)

var err error
var session *mgo.Session
var c *mgo.Collection

func connect() {
	log.Printf(
		"Will connect to local mongodb",
	)
	session, err = mgo.Dial("mongodb://localhost:27017")
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
	log.Printf(
		"Session will be closed",
	)
	session.Close()
	log.Printf(
		"Session closed",
	)
}

func readall() Books {
	log.Printf(
		"Books will be read",
	)
	result := Books{}
	err := c.Find(bson.M{}).All(&result)
	if err != nil {
		log.Printf("No books found", err)
	}
	log.Printf(
		"All books read",
	)
	return result
}

func read(name string) Book {
	log.Printf(
		"Book will be read",
	)
	result := Book{}
	err := c.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		log.Printf("Book not found:", name, err)
	}
	log.Printf(
		"Book read",
	)
	return result
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

func update(b Book) {
	oldBook := read(b.Name)
	if "" == oldBook.Name {
		log.Printf("Book does not exist:", b.Name)
	} else {
		err := c.Update(oldBook, b)
		if err != nil {
			panic(err)
		}
	}
}

func delete(name string) {
	log.Printf(
		"Book will be deleted",
	)
	err := c.Remove(bson.M{"name": name})
	if err != nil {
		log.Printf("Book not deleted:", name, err)
	}
	log.Printf(
		"Book deleted",
	)
}
