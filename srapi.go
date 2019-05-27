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
	srv.HandleFunc("/", getDummy).Methods("GET")
	srv.HandleFunc("/{sport}", getSport).Methods("GET")
	srv.HandleFunc("/{sport}/standings", getStandings).Methods("GET")
	return srv
}

func getDummy(w http.ResponseWriter, r *http.Request) {
	out, _ := json.Marshal("SRAPI")
	sendResponse(200, out, w)
	return
}

func getSport(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sport, err := queryDb(params["sport"])
	if err != nil {
		out, _ := json.Marshal("Sport not found.")
		sendResponse(500, out, w)
	} else {
		out, _ := json.Marshal(sport)
		sendResponse(200, out, w)
	}
	return
}

func getStandings(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sport, err := queryDb(params["sport"])
	if err != nil {
		out, _ := json.Marshal("Sport not found.")
		sendResponse(500, out, w)
		return
	}
	stand, _ := bsStandings(sport)
	out, err := json.Marshal(stand)
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
