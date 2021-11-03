package serveMux

import (
	"crypto/rand"
	"fmt"
	"net/http"
)

type UUID struct{}

func (p *UUID) ServeHTTP(w http.ResponseWriter, req http.Request) {
	if req.URL.Path == "/" {
		giveRandomUUID(w, req)
		return
	}
	
	http.NotFound(w, &req)
}

func giveRandomUUID(w http.ResponseWriter, req http.Request) {
	c := 10
	b := make([]byte, c)

	_, err := rand.Read(b)
	if err != nil{
		panic(err)
	}

	fmt.Fprintf(w, fmt.Sprintf("%x", b))
}