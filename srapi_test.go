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

	res = startEndpoint("/bs/standings/1997")
	assert.Equal(t, 200, res.Code, "Expected OK response")
	assert.Contains(t, res.Body.String(), "AL",
		"Expected result contains 'AL'")
	assert.Contains(t, res.Body.String(), "NL",
		"Expected result contains 'NL'")
	assert.Contains(t, res.Body.String(), "Central",
		"Expected result contains 'Central'")
	assert.Contains(t, res.Body.String(), "Florida Marlins",
		"Expected result contains 'Florida Marlins'")
	assert.NotContains(t, res.Body.String(), "Arizona Diamondbacks",
		"Expected result does not contain 'Arizona Diamondbacks'")
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

// Test an invalid endpoint for failure
func TestInvalidEndpoint(t *testing.T) {
	res := startEndpoint("/fakeUrl")
	assert.Equal(t, 404, res.Code, "Expected 404 response")
	assert.Contains(t, res.Body.String(), "page not found",
		"Expected proper error message")

	res = startEndpoint("/fakeUrl/partTwo")
	assert.Equal(t, 404, res.Code, "Expected 404 response")
}

// Test an invalid Sport ID for failure
func TestInvalidSport(t *testing.T) {
	res := startEndpoint("/xx/standings")
	assert.Equal(t, 500, res.Code, "Expected server error response")
	assert.Contains(t, res.Body.String(), "sql: no rows in result set",
		"Expected proper error message")
}
