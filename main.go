package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Start server at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", startServer()))
}
