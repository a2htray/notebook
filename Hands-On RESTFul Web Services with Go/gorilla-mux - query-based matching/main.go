package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		category := r.URL.Query().Get("category")

		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(w, "article in category %s", category)

	})

	log.Fatal(http.ListenAndServe(":8000", r))
}
