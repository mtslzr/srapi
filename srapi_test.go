package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDummy(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()
	startServer().ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code, "Expected OK response")
	assert.Equal(t, "SRAPI", res.Body.String(), "Expected 'SRAPI'")
}

func TestGetSport(t *testing.T) {
	req, _ := http.NewRequest("GET", "/bs", nil)
	res := httptest.NewRecorder()
	startServer().ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code, "Expected OK response")
	assert.Contains(t, res.Body.String(), "Baseball", "Expected result contains 'Baseball")
}

func TestGetSportFail(t *testing.T) {
	req, _ := http.NewRequest("GET", "/sb", nil)
	res := httptest.NewRecorder()
	startServer().ServeHTTP(res, req)

	assert.Equal(t, 500, res.Code, "Expected error response")
	assert.Equal(t, "Sport not found.", res.Body.String(), "Expected 'Sport not found.'")
}
