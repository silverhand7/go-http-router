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

type LogMiddleware struct {
	http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request Received")
	middleware.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Hello world http router")
	})
	middleware := LogMiddleware{
		router,
	}

	request := httptest.NewRequest("GET", "/", nil)
	recorder := httptest.NewRecorder()

	middleware.ServeHTTP(recorder, request)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	assert.Equal(t, "Hello world http router", string(body))
}
