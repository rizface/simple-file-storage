package test

import (
	"encoding/json"
	"file-storage/app/setup"
	"file-storage/model"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFileUpload(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		content,_ := ioutil.ReadFile("img.txt")
		request := model.FileRequest{
			Files: []string{
				string(content),
			},
		}
		requestJson,_ := json.Marshal(request)
		reader := strings.NewReader(string(requestJson))
		httpRequest := httptest.NewRequest(http.MethodPost,"http://localhost:8081/api/save",reader)
		recorder := httptest.NewRecorder()
		router := setup.Setup()
		router.ServeHTTP(recorder,httpRequest)
		resBody,_ := io.ReadAll(recorder.Body)
		fmt.Println(string(resBody))
	})
	t.Run("invalid", func(t *testing.T) {
		content,_ := ioutil.ReadFile("invalid-img.txt")
		request := model.FileRequest{
			Files: []string{
				string(content),
			},
		}
		requestJson,_ := json.Marshal(request)
		reader := strings.NewReader(string(requestJson))
		httpRequest := httptest.NewRequest(http.MethodPost,"http://localhost:8081/api/save",reader)
		recorder := httptest.NewRecorder()
		router := setup.Setup()
		router.ServeHTTP(recorder,httpRequest)
		resBody,_ := io.ReadAll(recorder.Body)
		fmt.Println(string(resBody))
		assert.Equal(t, http.StatusBadRequest,recorder.Code)
	})
}
