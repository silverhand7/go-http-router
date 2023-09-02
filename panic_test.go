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

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		fmt.Fprint(w, "Panic: ", i)
	}

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		panic("oops")
	})

	request := httptest.NewRequest("GET", "/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	assert.Equal(t, "Panic: oops", string(body))
}
