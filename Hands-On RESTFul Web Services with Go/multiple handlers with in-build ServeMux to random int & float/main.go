package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	newMux := http.NewServeMux()
	newMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, rand.Float64())
	})

	newMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, rand.Int())
	})

	if err := http.ListenAndServe(":8000", newMux); err != nil {
		log.Fatal(err)
	}
}
