package controller

import (
	"encoding/json"
	"file-storage/model"
	"file-storage/service"
	"net/http"
)

type fileControllerImpl struct{
	service service.FileService
}

func NewFileController(service service.FileService) FileController {
	return fileControllerImpl{service:service}
}

func (f fileControllerImpl) Save(w http.ResponseWriter, r *http.Request) {
	request := model.FileRequest{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&request)
	result := f.service.Save(request.Files)
	json.NewEncoder(w).Encode(model.Response{
		Code:   http.StatusOK,
		Status: "upload success",
		Data:   result,
	})
}

