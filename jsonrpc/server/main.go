package main

import (
	jsonparse "encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

type Args struct {
	ID string
}

type Book struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
}

type JSONServer struct{}

func (t *JSONServer) GiveBookDetail(req *http.Request, args *Args, reply *Book) error {
	var books []Book
	// read json file and load data
	absPath, _ := filepath.Abs("jsonrpc/server/books.json")
	raw, readerr:= ioutil.ReadFile(absPath)
	if readerr != nil {
		log.Println("error:", readerr)
		os.Exit(1)
	}

	// unmaarshal json data into books array 
	marshalerr := jsonparse.Unmarshal(raw, &books)
	if marshalerr != nil {
		log.Println("error: ", marshalerr)
		os.Exit(1)
	}

	// iterate over all the books to find the given book 
	for _, book := range books {
		if book.ID == args.ID {
		// if book found, fill reply with it 
		*reply = book
		break
		}
	}

	return nil
}

func main() {
	// create a new rpc server 
	s := rpc.NewServer()
	// register the type of data required as json
	s.RegisterCodec(json.NewCodec(), "application/json")

	// register the server by creating a new json server 
	s.RegisterService(new(JSONServer), "")

	mux := mux.NewRouter()
	mux.Handle("/rpc", s)

	http.ListenAndServe(":1234", mux)
}