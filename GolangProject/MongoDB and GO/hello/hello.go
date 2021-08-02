package main

import("log"
       "net/http"
       "github.com/gorilla/mux"
       "gopkg.in/go-playground/validator.v9"
      )

var validate *validator.Validate

func main() {
	validate = validator.New()
	//validate.RegisterValidation("age-under-18", ValidateAge)

	log.Print("[DEBUG] Starting server...")
	
	r := mux.NewRouter()
	r.HandleFunc("/signin", InsertUser).Methods("POST")
	r.HandleFunc("/login", LogInUser).Methods("POST")
	r.HandleFunc("/searchUser", SearchUser).Methods("GET")
	r.HandleFunc("/removeUser", RemoveUser).Methods("DELETE")
	r.HandleFunc("/updateUser", UpdateUser).Methods("PUT")
	r.HandleFunc("/allUsers", AllUsers).Methods("GET")

	r.HandleFunc("/createsticker", InsertSticker).Methods("POST")

	http.ListenAndServe("localhost:8080", r)
}