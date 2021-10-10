package controller

import "net/http"

type FileController interface {
	Save(w http.ResponseWriter, r *http.Request)
}
