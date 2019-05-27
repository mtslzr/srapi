package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Start server and query and endpoint.
func startEndpoint(ep string) (res *httptest.ResponseRecorder) {
	req, _ := http.NewRequest("GET", ep, nil)
	res = httptest.NewRecorder()
	startServer().ServeHTTP(res, req)
	return
}

// Test GetStandings for current baseball standings.
func TestGetStandings(t *testing.T) {
	res := startEndpoint("/bs/standings")
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

// Test GetTeams for current baseball teams.
func TestGetTeams(t *testing.T) {
	res := startEndpoint("/bs/teams")
	assert.Equal(t, 200, res.Code, "Expected OK response")
	assert.Contains(t, res.Body.String(), "Miami Marlins",
		"Expected result contains 'Miami Marlins'")
	assert.Contains(t, res.Body.String(), "Minnesota Twins",
		"Expected result contains 'Minnesota Twins'")
}

// Test GetTeams for current baseball teams.
func TestGetYears(t *testing.T) {
	res := startEndpoint("/bs/years")
	assert.Equal(t, 200, res.Code, "Expected OK response")
	assert.Contains(t, res.Body.String(), "2019",
		"Expected result contains '2019'")
	assert.Contains(t, res.Body.String(), "1876",
		"Expected result contains '1876'")
}
