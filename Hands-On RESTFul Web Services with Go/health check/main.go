package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now()
		_, _= io.WriteString(w, currentTime.String())
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
