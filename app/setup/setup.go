package setup

import (
	"file-storage/app"
	"file-storage/app/middleware"
	"file-storage/controller"
	"file-storage/service"
	"github.com/gorilla/mux"
	"net/http"
)

func Setup() *mux.Router {
	router := app.Mux.NewRoute().Subrouter()
	router.Use(middleware.PanicHandler)

	fileService := service.NewFileServiceImpl()
	fileController := controller.NewFileController(fileService)

	router.HandleFunc(app.SAVE,fileController.Save).Methods(http.MethodPost)
	return router
}
