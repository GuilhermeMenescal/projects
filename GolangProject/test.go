package main

import("fmt"
       "log"
       "net/http"
       "github.com/gorilla/mux"
       "gopkg.in/mgo.v2"
    //  "gopkg.in/mgo.v2/bson"
	)
type User struct{
      Username string
      Email string
      Password string
      Fullname string	
}

func main() {

	fmt.Print("Bem vindo!")
	r := mux.NewRouter()
	r.HandleFunc("/register", InsertUser).Methods("POST")

	http.ListenAndServe("localhost:8080", r)
}


func InsertUser(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial("localhost:8080")
	if err != nil {
                panic(err)
        }
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("my_database").C("users")
	err = c.Insert(&User{"Ale", "banana@bananinha", "banana", "Banale"},
	               &User{"Cla", "clanana@clananinha", "banana", "Clanana"})
	 if err != nil {
                log.Fatal( fmt.Print("erro"))
        }
	fmt.Print("Usuário registrado")
	return w.Write("reeeeeee")
}
