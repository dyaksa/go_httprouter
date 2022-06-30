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

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, response *http.Request) {
	fmt.Print("before Middleware")
	middleware.Handler.ServeHTTP(writer, response)
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "middleware")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	response := httptest.NewRecorder()

	logMiddleware := LogMiddleware{
		Handler: router,
	}

	logMiddleware.ServeHTTP(response, request)

	body, _ := io.ReadAll(response.Result().Body)
	assert.Equal(t, "middleware", string(body))
}
