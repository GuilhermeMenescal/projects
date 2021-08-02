package config

import ("net/http"
       "github.com/gorilla/mux")

func StartServer() {
	r := mux.NewRouter()
	SetupUsersRoutes(r)
	http.ListenAndServe(SERVER_HOST, r)
}
