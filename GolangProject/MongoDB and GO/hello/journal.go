package main

import ("encoding/json"
	"net/http"
	"log"
	//"gopkg.in/mgo.v2"
	)


type Journal struct {
	
	Title string     `json:Title validate:"required"`
	//User User 		 `validate:"required"`
	//Pages int32      `validate:"required", "lt=365"`
}

func InsertJournal (w http.ResponseWriter, r *http.Request){
	body := r.Body
	d    := json.NewDecoder(body)

	var j Journal	

	err := d.Decode(&j)
	
	if err != nil {
		log.Print("[ERROR] wrong JSON")	
	}
	w.Write([]byte("InsertJournal TO-DO"))
return
}