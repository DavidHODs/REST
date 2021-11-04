package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/justinas/alice"
)

type city struct {
	Name string
	Area uint64
}

func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, req *http.Request) {
		log.Println("Currently in the check content type middleware")

		// filtering requests by MIME type 
		if req.Header.Get("Content-type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsupported media type. Send JSON"))
			return
		}

		handler.ServeHTTP(w, req)
	})
}

func setServerTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, req *http.Request) {
		handler.ServeHTTP(w, req)

		// setting cookie to every api response 
		cookie := http.Cookie{Name: "Server-Time(UTC)", Value: strconv.FormatInt(time.Now().Unix(), 10)}

		http.SetCookie(w, &cookie)
		log.Println("Currently in the set cookie middleware")
	})
}

func postHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		var tempCity city
		err := json.NewDecoder(req.Body).Decode(&tempCity)
		if err != nil {
			panic(err)
		}
		defer req.Body.Close()

		fmt.Printf("Got %s city with area of %d sq miles\n", tempCity.Name, tempCity.Area)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("201-created"))	
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405-method not allowed"))
	}
}

func main() {
	// http.HandleFunc("/city", postHandler)
	originalHandler := http.HandlerFunc(postHandler)
	// http.Handle("/city", filterContentType(setServerTimeCookie(originalHandler)))
	chain := alice.New(filterContentType, setServerTimeCookie).Then(originalHandler)
	http.Handle("/city", chain)
	http.ListenAndServe(":8080", nil)
}