package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Start server at localhost:5000...")
	http.ListenAndServe(":5000", startServer())
}
