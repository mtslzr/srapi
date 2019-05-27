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
	assert.Contains(t, res.Body.String(), "SRAPI",
		"Expected 'SRAPI'")
}

func TestGetSport(t *testing.T) {
	req, _ := http.NewRequest("GET", "/bs", nil)
	res := httptest.NewRecorder()
	startServer().ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code, "Expected OK response")
	assert.Contains(t, res.Body.String(), "Baseball",
		"Expected result contains 'Baseball")
}

func TestGetSportFail(t *testing.T) {
	req, _ := http.NewRequest("GET", "/sb", nil)
	res := httptest.NewRecorder()
	startServer().ServeHTTP(res, req)

	assert.Equal(t, 500, res.Code, "Expected error response")
	assert.Contains(t, res.Body.String(), "Sport not found.",
		"Expected 'Sport not found.'")
}

func TestGetStandings(t *testing.T) {
	req, _ := http.NewRequest("GET", "/bs/standings", nil)
	res := httptest.NewRecorder()
	startServer().ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code, "Expected OK response")
	assert.Contains(t, res.Body.String(), "AL",
		"Expected result contains 'AL'")
	assert.Contains(t, res.Body.String(), "NL",
		"Expected result contains 'NL'")
	assert.Contains(t, res.Body.String(), "West",
		"Expected result contains 'West'")
	assert.Contains(t, res.Body.String(), "Miami Marlins",
		"Expected result contains 'Miami Marlins'")
	assert.Contains(t, res.Body.String(), "Minnesota Twins",
		"Expected result contains 'Minnesota Twins'")
}
