package app

import "github.com/gorilla/mux"

var Mux *mux.Router = mux.NewRouter()

const(
	HOST = "http://localhost:8081/my/"
	SAVE = "/api/save"
)
