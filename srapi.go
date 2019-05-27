package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// Sport is model for Database data
type Sport struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Host      string `json:"host"`
	Standings string `json:"standings"`
	Teams     string `json:"teams"`
	Years     string `json:"years"`
}

func startServer() http.Handler {
	srv := mux.NewRouter()
	srv.HandleFunc("/{sport}/standings", getStandings).Methods("GET")
	srv.HandleFunc("/{sport}/teams", getTeams).Methods("GET")
	return srv
}

func getSport(id string, w http.ResponseWriter) Sport {
	sport, err := queryDb(id)
	if err != nil {
		out, _ := json.Marshal(err)
		sendResponse(500, out, w)
	}
	return sport
}

func getStandings(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sport := getSport(params["sport"], w)

	stand, err := bsStandings(sport)
	if err != nil {
		out, _ := json.Marshal(err)
		sendResponse(500, out, w)
	}
	out, _ := json.Marshal(stand)
	sendResponse(200, out, w)
	return
}

func getTeams(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sport := getSport(params["sport"], w)

	teams, err := bsTeams(sport)
	if err != nil {
		out, _ := json.Marshal(err)
		sendResponse(500, out, w)
	}
	out, _ := json.Marshal(teams)
	sendResponse(200, out, w)
	return
}

func queryDb(sport string) (row Sport, err error) {
	db, err := sql.Open("sqlite3", "./srapi.db")
	if err != nil {
		return row, err
	}
	err = db.QueryRow("SELECT * FROM sports WHERE id = ?", sport).Scan(&row.ID, &row.Name, &row.Host, &row.Standings, &row.Teams, &row.Years)
	if err != nil {
		return row, err
	}
	return row, nil
}

func sendResponse(code int, js []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(js)
	return
}
