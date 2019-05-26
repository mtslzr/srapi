package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func startServer() http.Handler {
	srv := mux.NewRouter()
	srv.HandleFunc("/", getDummy)
	return srv
}

func getDummy(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("SRAPI")
	return
}
