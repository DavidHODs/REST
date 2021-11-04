package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct{}

type TimeServer int64

func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	// fill reply pointer to send the data back

	*reply = time.Now().Unix()
	return nil
}

func main() {
	timeServer := new(TimeServer)
	rpc.Register(timeServer)
	rpc.HandleHTTP()

	// listen to requests on port 1234 
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen Error:", err)
	}

	http.Serve(listen, nil)
}