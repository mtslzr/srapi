package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Start server at localhost:5000...")
	err := http.ListenAndServe(":5000", startServer())
	if err != nil {
		log.Fatal(err)
	}
}
