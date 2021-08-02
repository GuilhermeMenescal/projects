package main

import ("encoding/json"
	"net/http"
	"log"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	)


type User struct {
	
	Fullname string     `json:fullname validate:"required, alphaunicode"`
	Username string `json:username validate:"required, alphanumunicode"`
	Email string    `json:email validate:"required,email"`
	Password string `json:pass validate:"required"`
	Age int         `json:age validate:"required"`
	Phone string    `json:phone validate: "required, len=12, numeric"`
	Followers int64
	Following int64
}

var m map[string]User

func InsertUser(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	d    := json.NewDecoder(body)

	var u User	

	err := d.Decode(&u)
	
	if err != nil {
		log.Print("[ERROR] wrong JSON")	
	}

	err = validate.Struct(u)

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

	c := session.DB("db_bujo").C("users")
        err = c.Insert(&u)
        if err != nil {
               log.Print("[ERROR] Database failled")	
        }


	w.Write([]byte(u.Fullname + "     " + u.Email + u.Password))
	return

}

func SearchUser(w http.ResponseWriter, r *http.Request){
m = make(map[string]User)

w.Write([]byte("SearchUsers TO-DO"))
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


	w.Write([]byte("AllUsers TO-DO"))
	return

}

func LogInUser(w http.ResponseWriter, r *http.Request){

w.Write([]byte("LogIn TO-DO"))
return
}

func RemoveUser(w http.ResponseWriter, r *http.Request){

w.Write([]byte("RemoveUser TO-DO"))
return
}



func UpdateUser(w http.ResponseWriter, r *http.Request){

w.Write([]byte("UpdateUser TO-DO"))
return
}
