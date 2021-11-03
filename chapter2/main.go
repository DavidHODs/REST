package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	err := http.ListenAndServe(":8080", mux)

	log.Fatal(err)
}