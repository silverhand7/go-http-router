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

func TestNotFound(t *testing.T) {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Not found")
	})

	request := httptest.NewRequest("GET", "/not-found", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	assert.Equal(t, "Not found", string(body))
}
