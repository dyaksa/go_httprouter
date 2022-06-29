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

func TestNameParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:id/item/:itemId", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		itemId := params.ByName("itemId")
		text := fmt.Sprintf("product %s itemId %s", id, itemId)
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/product/1/item/1", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	body, _ := io.ReadAll(response.Result().Body)
	assert.Equal(t, "product 1 itemId 1", string(body))
}

func TestPatternCatchAll(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*images", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		pattern := params.ByName("images")
		text := fmt.Sprintf("images : %s", pattern)
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/images/source/dias.png", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	body, _ := io.ReadAll(response.Result().Body)

	assert.Equal(t, "images : /source/dias.png", string(body))
}
