package helper

import (
	"encoding/json"
	"file-storage/model"
	"net/http"
)

func JsonWriter(w http.ResponseWriter, code int, status string ,data []string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(model.Response{
		Code:   code,
		Status: status,
		Data:   data,
	})
}
