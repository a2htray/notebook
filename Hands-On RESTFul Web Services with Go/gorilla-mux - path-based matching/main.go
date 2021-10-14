package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", articleHandler)

	srv := &http.Server{
		Handler: r,
		Addr: ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func articleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	_, _ = fmt.Fprintf(w, "Category is: %v\n", vars["category"])
	_, _ = fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}