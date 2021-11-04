package main

import (
	"fmt"
	"net/http"
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

func main() {
	originalHandler := http.HandlerFunc(handle)
	http.Handle("/", middleware(originalHandler))

	http.ListenAndServe(":8080", nil)
}