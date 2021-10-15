package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func articleHandler(w http.ResponseWriter, r *http.Request) {

}

func dynamicRoute(r *mux.Router) {
	url, err := r.Get("articleRoute").URL("category", "books", "id", "123")
	if err != nil {
		panic(err)
	}

	fmt.Printf("article url generated dynamically: %s\n", url.String())
}

func main() {
	r := mux.NewRouter()

	r.StrictSlash(false)
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", articleHandler).Name("articleRoute")

	dynamicRoute(r)

	log.Fatal(http.ListenAndServe(":8000", r))
}
