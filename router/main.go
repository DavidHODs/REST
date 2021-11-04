package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func ArticleHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "category is: %v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}

func QueryHandler(w http.ResponseWriter, req *http.Request) {
	queryParams := req.URL.Query()
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Got parameter id: %s\n", queryParams["id"][0])
	fmt.Fprintf(w, "Got parameter category: %s\n", queryParams["category"][0])
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
	mux.HandleFunc("/articles", QueryHandler)

	port := ":8080"

	srv := &http.Server{
		Handler: mux,
		Addr: port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	err := srv.ListenAndServe()

	fmt.Printf("Listening on port %s", port)

	log.Fatal(err)
}