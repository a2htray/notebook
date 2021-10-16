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

func (t *TimeServer) GiveServerTime(args *Args, replay *int64) error {
	*replay = time.Now().Unix()
	return nil
}

func main() {
	timeServer := new(TimeServer)

	rpcServer := rpc.NewServer()
	_ = rpcServer.Register(timeServer)
	rpcServer.HandleHTTP("./go/rpc", "./go/rpc/debug")

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal(e)
	}

	_ = http.Serve(l, nil)
}
