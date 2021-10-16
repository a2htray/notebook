package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s url: %s\n", time.Now().Format(time.RFC3339), r.URL.String())
		handler.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/book", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "book")
	})

	mux.HandleFunc("/api/article", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "article")
	})

	server := &http.Server{
		Addr:         ":8000",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Handler:      middleware(mux),
	}

	log.Fatal(server.ListenAndServe())
}
