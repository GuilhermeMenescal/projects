package main

import("log"
	"main/config")


func main() {
	log.Print("[DEBUG] Starting server...")
	
	err := config.CreateDBConnection()

	if err != nil {
		panic(err)
	}

	config.StartValidator()
	config.StartServer()
}





















