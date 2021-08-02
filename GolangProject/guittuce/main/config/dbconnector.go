package config

import("gopkg.in/mgo.v2"
		"log")

var session *mgo.Session


func CreateDBConnection() (err error){

	session, err = mgo.Dial(MONGO_HOST)

	if err != nil {
		log.Print("[ERROR] failed to connect to database")
		return 
	}

	return
}

func GetSession() *mgo.Session {

	return session
}

func OpenSession(collection string) *mgo.Collection {
	return session.DB("db_guittuce").C(collection)
}