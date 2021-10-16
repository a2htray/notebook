package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	type Args struct{}

	var replay int64
	args := Args{}

	// client, err := rpc.DialHTTP("tcp", "localhost:1234")

	client, err := rpc.DialHTTPPath("tcp", "localhost:1234", "./go/rpc")
	if err != nil {
		log.Fatal(err)
	}

	err = client.Call("TimeServer.GiveServerTime", &args, &replay)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("replay: %d\n", replay)
}
