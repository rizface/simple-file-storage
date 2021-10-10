package main

import (
	"file-storage/app"
	"file-storage/app/setup"
	"net/http"
)

func main() {
	//router := mux.NewRouter()
	//router.Use(middleware.PanicHandler)
	//
	//fileService := service.NewFileServiceImpl()
	//fileController := controller.NewFileController(fileService)
	//
	//router.HandleFunc(app.SAVE,fileController.Save).Methods(http.MethodPost)
	setup.Setup()
	server := http.Server{
		Addr: ":8081",
		Handler: app.Mux,
	}
	fileServer := http.FileServer(http.Dir("files"))
	app.Mux.PathPrefix("/my/").Handler(http.StripPrefix("/my/",fileServer))
	server.ListenAndServe()
}
