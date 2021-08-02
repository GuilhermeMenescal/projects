package main

import ("encoding/json"
	"net/http"
	"log"
	"gopkg.in/mgo.v2"
	)


type Sticker struct {
	
	Name string     `json:name validate:"required"`
	Url string		 `json:url validate: "required"`
}

func InsertSticker(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	d    := json.NewDecoder(body)

	var s Sticker	

	err := d.Decode(&s)
	
	if err != nil {
		log.Print("[ERROR] wrong JSON")	
	}

	err = validate.Struct(s)

	if err != nil {
		log.Printf("[ERROR] Wrong data: %v", err)
		w.Write([]byte(`{"code": 400, "message":"Invalid informations"}`))
		return	
	}


	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("db_bujo").C("stickers")
        err = c.Insert(&s)
        if err != nil {
               log.Print("[ERROR] Database failled")	
        }


	w.Write([]byte(s.Name + " " + s.Url))
	return

}

func SearchUser(w http.ResponseWriter, r *http.Request){
m = make(map[string]Sticker)

w.Write([]byte("SearchSticker TO-DO"))
return
}

func AllUsers(w http.ResponseWriter, r *http.Request){

//	session, err := mgo.Dial("localhost")
//	if err != nil {
//		panic(err)
//	}
//	defer session.Close()
//	session.SetMode(mgo.Monotonic, true)

//	c := session.DB("db_bujo").C("users")
 //       err = c.Find(bson.M{})
   //     if err != nil {
     //          log.Print("[ERROR] Database failled")	
       // }


	w.Write([]byte("AllStickers TO-DO"))
	return

}