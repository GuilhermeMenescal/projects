package main

import ("encoding/json"
	"net/http"
	"log"
	"gopkg.in/mgo.v2")


type User struct {
	
	Name string     `json:name`
	Email string    `json:email`
	Password string `json:pass`
	Age int         `json:age`
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	d    := json.NewDecoder(body)

	var u User	

	err := d.Decode(&u)
	
	if err != nil {
		log.Print("[ERROR] wrong JSON")	
	}

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("db_bujo").C("users")
        err = c.Insert(&u)
        if err != nil {
               log.Print("[ERROR] Database failled")	
        }


	w.Write([]byte(u.Name + "     " + u.Email + u.Password))
	return

}