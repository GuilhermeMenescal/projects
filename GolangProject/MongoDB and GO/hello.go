package main

import("log"
       "net/http"
       "github.com/gorilla/mux"
      )


func main() {

	log.Print("[DEBUG] Starting server...")
	
	r := mux.NewRouter()
	r.HandleFunc("/users", InsertUser).Methods("POST")

	http.ListenAndServe("localhost:8080", r)
}