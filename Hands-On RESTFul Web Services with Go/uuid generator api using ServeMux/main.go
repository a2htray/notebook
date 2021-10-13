package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
)

type UUID struct {

}

func (u *UUID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		c := 10
		b := make([]byte, c)
		_, err := rand.Read(b)
		if err != nil {
			panic(err)
		}

		_, _ = fmt.Fprintf(w, fmt.Sprintf("%x", b))
	} else {
		http.NotFound(w, r)
	}
	return
}

func main() {
	mux := &UUID{}
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}
}
