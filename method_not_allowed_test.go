package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMethodNotAllowed(t *testing.T) {
	router := httprouter.New()

	router.MethodNotAllowed = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Tidak Diijinkan")
	})

	router.POST("/upload", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "upload form file")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/upload", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	body, _ := io.ReadAll(response.Result().Body)

	assert.Equal(t, "Tidak Diijinkan", string(body))
}
