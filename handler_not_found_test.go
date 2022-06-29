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

func TestHandlerNotFound(t *testing.T) {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Halaman Tidak Ketemu")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	body, _ := io.ReadAll(response.Result().Body)

	assert.Equal(t, "Halaman Tidak Ketemu", string(body))
}
