package main

import (
	"encoding/json"
	"github.com/justinas/alice"
	"log"
	"net/http"
	"strconv"
	"time"
)

func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Currently in the filter content type middleware")
		if r.Header.Get("content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			_, _ = w.Write([]byte("415 - Unsupported Media Type, Please send JSON"))
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func setServerTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		cookie := http.Cookie{
			Name:     "Server Time(UTC)",
			Value:    strconv.FormatInt(time.Now().Unix(), 10),
			HttpOnly: true,
			Secure:   false,
		}

		http.SetCookie(w, &cookie)
		log.Println("Currently in the set server time middleware")
	})
}

type City struct {
	Name string `json:"name"`
	Area int64  `json:"area"`
}

var cities = make(map[string]*City)

func postSaveCity(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		city := City{}
		decoder := json.NewDecoder(r.Body)
		_ = decoder.Decode(&city)
		cities[city.Name] = &city

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte("201 - Created"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("405 - Method Not Allowed, Please use POST"))
	}
}

func getListCities(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data := make([]*City, len(cities))
		counter := 0
		for _, city := range cities {
			data[counter] = city
			counter++
		}

		w.Header().Set("content-Type", "application/json")
		dataBytes, _ := json.Marshal(data)
		_, _ = w.Write(dataBytes)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("405 - Method Not Allowed, Please use GET"))
	}
}

func main() {
	postCityChan := alice.New(filterContentType, setServerTimeCookie).ThenFunc(postSaveCity)
	getCitiesChan := alice.New(setServerTimeCookie).ThenFunc(getListCities)
	http.Handle("/api/city", postCityChan)
	http.Handle("/api/cities", getCitiesChan)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
