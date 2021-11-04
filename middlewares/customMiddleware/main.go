package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func middleware(originalHandler http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, req *http.Request) {
		fmt.Println("Executing middleware before request phase")

		// pass control back to the handler 
		originalHandler.ServeHTTP(w, req)
		fmt.Println("Executing middleware after response phase")
	})
}

func handle(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Executing main handler")
	w.Write([]byte("OK"))
}

func handleLogger(w http.ResponseWriter, req *http.Request) {
	log.Println("Processing Request")
	w.Write([]byte("OK"))
	log.Println("Finished processing request")
}

func main() {
	// originalHandler := http.HandlerFunc(handle)
	// http.Handle("/", middleware(originalHandler))

	// http.ListenAndServe(":8080", nil)

	mux := mux.NewRouter()
	mux.HandleFunc("/", handleLogger)
	loggedRouter := handlers.LoggingHandler(os.Stdout, mux)

	err := http.ListenAndServe(":8080", loggedRouter)
	log.Fatal(err)
}