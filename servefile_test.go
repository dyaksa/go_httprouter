package main

import (
	"embed"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()
	directory, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(directory))

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/files/hello.txt", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	body, _ := io.ReadAll(response.Result().Body)
	assert.Equal(t, "Hello Resources", string(body))

}
