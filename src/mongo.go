package main

import (
	mgo "gopkg.in/mgo.v2"
)

type Mongo struct {
	DB       *mgo.Database
	PostPool *mgo.Collection
}

func (m *Mongo) login() {
	m.DB = m.getDB()
	// m.DB.Login("", "")
	m.PostPool = m.DB.C("post")
}

func (m *Mongo) savePost(post Post) {
	err := m.PostPool.Insert(&post)
	if err != nil {
		panic(err)
	}
}

func (m *Mongo) getDB() *mgo.Database {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	db := session.DB("dtalkGO") //root user is created in the admin authentication database and given the role of root.
	return db
}
