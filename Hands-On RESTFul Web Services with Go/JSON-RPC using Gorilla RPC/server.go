package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	rpcjson "github.com/gorilla/rpc/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Book struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
}

type Args struct {
	ID string
}

type JSONServer int64

func (j *JSONServer) GiveBookDetail(r *http.Request, args *Args, replay *Book) error {
	var books []Book
	jsonFilepath, _ := filepath.Abs("./books.json")
	dataOfBooks, err := ioutil.ReadFile(jsonFilepath)
	if err != nil {
		log.Println("book: read file error")
		os.Exit(-1)
	}

	if err = json.Unmarshal(dataOfBooks, &books); err != nil {
		log.Println("book: parse json data error")
		os.Exit(-1)
	}

	for _, book := range books {
		if book.ID == args.ID {
			*replay = book
			break
		}
	}
	return nil
}

func main() {
	rpcServer := rpc.NewServer()
	rpcServer.RegisterCodec(rpcjson.NewCodec(), "application/json")

	err := rpcServer.RegisterService(new(JSONServer), "")
	if err != nil {
		log.Println("rpc: register service error - ", err.Error())
		os.Exit(-1)
	}

	r := mux.NewRouter()
	r.Handle("/rpc", rpcServer)

	log.Fatal(http.ListenAndServe(":1234", r))
}
