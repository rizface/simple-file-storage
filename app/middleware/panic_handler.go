package middleware

import (
	"file-storage/app/exception"
	"file-storage/helper"
	"net/http"
)

func PanicHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if badRequest, badRequestOK := err.(exception.BadRequest); badRequestOK {
				helper.JsonWriter(w,http.StatusBadRequest,badRequest.Error.(string),[]string{})
			}
		}()
		next.ServeHTTP(w,r)
	})
}
