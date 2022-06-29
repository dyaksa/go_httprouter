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

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()
	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, error interface{}) {
		fmt.Fprint(writer, "Panic : ", error)
	}

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		panic("error")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	body, _ := io.ReadAll(response.Result().Body)
	assert.Equal(t, "Panic : error", string(body))
}
