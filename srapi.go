package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var w http.ResponseWriter

// Sport is model for Database data
type Sport struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Host      string `json:"host"`
	Standings string `json:"standings"`
	Teams     string `json:"teams"`
	Years     string `json:"years"`
}

// Start API and set routes
func startServer() http.Handler {
	srv := mux.NewRouter()
	srv.HandleFunc("/{sport}/standings", getStandings).Methods("GET")
	srv.HandleFunc("/{sport}/teams", getTeams).Methods("GET")
	srv.HandleFunc("/{sport}/years", getYears).Methods("GET")
	srv.Use(rwMiddleware)
	return srv
}

// Get Sport by ID
func getSport(id string) Sport {
	sport, err := queryDb(id)
	checkError(err)
	return sport
}

// Get Current Standings
func getStandings(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sport := getSport(params["sport"])
	stand := bsStandings(sport)

	out, _ := json.Marshal(stand)
	sendResponse(200, out)
}

// Get All Teams
func getTeams(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sport := getSport(params["sport"])
	teams := bsTeams(sport)

	out, _ := json.Marshal(teams)
	sendResponse(200, out)
}

// Get All Years
func getYears(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sport := getSport(params["sport"])
	years := bsYears(sport)

	out, _ := json.Marshal(years)
	sendResponse(200, out)
}

// Query Database and return Sport
func queryDb(sport string) (row Sport, err error) {
	db, err := sql.Open("sqlite3", "./srapi.db")
	checkError(err)

	err = db.QueryRow("SELECT * FROM sports WHERE id = ?", sport).
		Scan(&row.ID, &row.Name, &row.Host, &row.Standings, &row.Teams, &row.Years)
	checkError(err)

	return row, nil
}

// Export http.ResponseWriter for use in sendResponse()
func rwMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		w = rw
		next.ServeHTTP(rw, r)
	})
}

// If error, send 500 response
func checkError(err error) {
	if err != nil {
		msg, _ := json.Marshal(err)
		sendResponse(500, msg)
	}
}

// Marshal and send response with chosen statusCode
func sendResponse(code int, js []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(js)
	checkError(err)
}
