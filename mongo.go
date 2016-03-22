package main

import (
	"gopkg.in/mgo.v2/mgo"
	"gopkg.in/mgo.v2/mgo/bson"
	"log"
)

type mongo struct{}

func connect(url string) (s *mgo.Session, err error) {
	s, err = mgo.Dial(url)
	if err != nil {
		return
	}
	log.Printf("Connected to: ", url)
	s.SetMode(mgo.Monotonic, true)
	return
}

func collection(session *mgo.Session, dbname, cname string) (c *mgo.Collection) {
	c = session.DB(dbname).C(cname)
	log.Printf("Session retireved")

	return
}

func close(s *mgo.Session) {
	s.Close()
	log.Printf("Session closed")
}

func (m *mongo) Readall(url, dbName, cName string) (books Books) {
	s, err := connect(url)
	if err != nil {
		panic(err)
	}
	c := collection(s, dbName, cName)
	defer close(s)

	log.Printf("Books will be read")
	result := Books{}
	err = c.Find(bson.M{}).All(&result)
	if err != nil {
		log.Printf("No books found", err)
	}
	log.Printf("All books read")
	return result
}

func (m *mongo) Read(url, dbName, cName, bookId string) Book {
	s, err := connect(url)
	if err != nil {
		panic(err)
	}
	c := collection(s, dbName, cName)
	defer close(s)

	log.Printf("Book will be read")
	result := Book{}
	err = c.Find(bson.M{"name": bookId}).One(&result)
	if err != nil {
		log.Printf("Book not found:", bookId, err)
	}
	log.Printf("Book read")
	return result
}

func (m *mongo) Write(url, dbName, cName string, b Book) {
	s, err := connect(url)
	if err != nil {
		panic(err)
	}
	c := collection(s, dbName, cName)
	defer close(s)

	oldBook := m.Read(url, dbName, cName, b.Name)
	if "" != oldBook.Name {
		log.Printf("Book already exist:", oldBook.Name)
	} else {
		err := c.Insert(b)
		if err != nil {
			panic(err)
		}
	}
}

func (m *mongo) Update(url, dbName, cName string, b Book) {
	s, err := connect(url)
	if err != nil {
		panic(err)
	}
	c := collection(s, dbName, cName)
	defer close(s)

	oldBook := m.Read(url, dbName, cName, b.Name)
	if "" == oldBook.Name {
		log.Printf("Book does not exist:", b.Name)
	} else {
		err := c.Update(oldBook, b)
		if err != nil {
			panic(err)
		}
	}
}

func (m *mongo) Delete(url, dbName, cName, bookId string) {
	s, err := connect(url)
	if err != nil {
		panic(err)
	}
	c := collection(s, dbName, cName)
	defer close(s)

	log.Printf("Book will be deleted")
	err = c.Remove(bson.M{"name": bookId})
	if err != nil {
		log.Printf("Book not deleted:", bookId, err)
	}
	log.Printf("Book deleted")
}
