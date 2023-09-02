package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouterNamedParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id/items/:itemId", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		itemId := p.ByName("itemId")
		text := "Product " + id + " item " + itemId
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest("GET", "/products/1/items/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	assert.Equal(t, "Product 1 item 1", string(body))
}

func TestPatternCatchAll(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		text := "Image: " + p.ByName("image")
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest("GET", "/images/small/profile.jpg", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	assert.Equal(t, "Image: /small/profile.jpg", string(body))
}
