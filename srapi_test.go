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
func TestGetStandingsCurrent(t *testing.T) {
	res := startEndpoint("/bs/standings")
	assert.Equal(t, 200, res.Code, "Expected OK response")
	assert.Contains(t, res.Body.String(), "AL",
		"Expected result contains 'AL'")
	assert.Contains(t, res.Body.String(), "NL",
		"Expected result contains 'NL'")
	assert.Contains(t, res.Body.String(), "East",
		"Expected result contains 'East'")
	assert.Contains(t, res.Body.String(), "Central",
		"Expected result contains 'Central'")
	assert.Contains(t, res.Body.String(), "West",
		"Expected result contains 'West'")
	assert.Contains(t, res.Body.String(), "pos",
		"Expected result contains 'pos'")
	assert.Contains(t, res.Body.String(), "abbr",
		"Expected result contains 'abbr'")
	assert.Contains(t, res.Body.String(), "Miami Marlins",
		"Expected result contains 'Miami Marlins'")
	assert.Contains(t, res.Body.String(), "Minnesota Twins",
		"Expected result contains 'Minnesota Twins'")
}

// Test GetStandings for previous year.
func TestGetStandingsYear(t *testing.T) {
	res := startEndpoint("/bs/standings/1997")
	assert.Equal(t, 200, res.Code, "Expected OK response")
	assert.Contains(t, res.Body.String(), "AL",
		"Expected result contains 'AL'")
	assert.Contains(t, res.Body.String(), "NL",
		"Expected result contains 'NL'")
	assert.Contains(t, res.Body.String(), "East",
		"Expected result contains 'East'")
	assert.Contains(t, res.Body.String(), "Central",
		"Expected result contains 'Central'")
	assert.Contains(t, res.Body.String(), "West",
		"Expected result contains 'West'")
	assert.Contains(t, res.Body.String(), "pos",
		"Expected result contains 'pos'")
	assert.Contains(t, res.Body.String(), "abbr",
		"Expected result contains 'abbr'")
	assert.Contains(t, res.Body.String(), "Florida Marlins",
		"Expected result contains 'Florida Marlins'")
	assert.NotContains(t, res.Body.String(), "Arizona Diamondbacks",
		"Expected result does not contain 'Arizona Diamondbacks'")
}

// Test GetStandings for previous year with fewer divisions.
func TestGetStandingsYearFewDivs(t *testing.T) {
	res := startEndpoint("/bs/standings/1974")
	assert.Equal(t, 200, res.Code, "Expected OK response")
	assert.Contains(t, res.Body.String(), "AL",
		"Expected result contains 'AL'")
	assert.Contains(t, res.Body.String(), "NL",
		"Expected result contains 'NL'")
	assert.Contains(t, res.Body.String(), "East",
		"Expected result contains 'East'")
	assert.Contains(t, res.Body.String(), "West",
		"Expected result contains 'West'")
	assert.Contains(t, res.Body.String(), "pos",
		"Expected result contains 'pos'")
	assert.Contains(t, res.Body.String(), "abbr",
		"Expected result contains 'abbr'")
	assert.Contains(t, res.Body.String(), "Montreal Expos",
		"Expected result contains 'Montreal Expos'")
	assert.NotContains(t, res.Body.String(), "Colorado Rockies",
		"Expected result does not contain 'Colorado Rockies'")
}

// Test GetStandings for previous year with no divisions.
func TestGetStandingsYearNoDiv(t *testing.T) {
	res := startEndpoint("/bs/standings/1919")
	assert.Equal(t, 200, res.Code, "Expected OK response")
	assert.Contains(t, res.Body.String(), "AL",
		"Expected result contains 'AL'")
	assert.Contains(t, res.Body.String(), "NL",
		"Expected result contains 'NL'")
	assert.NotContains(t, res.Body.String(), "East",
		"Expected result contains 'East'")
	assert.Contains(t, res.Body.String(), "pos",
		"Expected result contains 'pos'")
	assert.Contains(t, res.Body.String(), "abbr",
		"Expected result contains 'abbr'")
	assert.Contains(t, res.Body.String(), "Brooklyn Dodgers",
		"Expected result contains 'Brooklyn Dodgers'")
	assert.NotContains(t, res.Body.String(), "New York Mets",
		"Expected result does not contain 'New York Mets'")
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
