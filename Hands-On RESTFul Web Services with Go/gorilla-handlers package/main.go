package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func handle(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("your request had been handled"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/handle", handle)
	handler := handlers.LoggingHandler(os.Stdout, r)

	log.Fatal(http.ListenAndServe(":8000", handler))

}
