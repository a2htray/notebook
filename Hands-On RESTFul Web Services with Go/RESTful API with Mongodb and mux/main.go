package main

import (
	"context"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

type DB struct {
	collection *mongo.Collection
}

func (d DB) GetMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var movie Movie

	objectID, _ := primitive.ObjectIDFromHex(vars["id"])
}

func (d DB) PostMovie(w http.ResponseWriter, r *http.Request) {

}

type Movie struct {
	ID        interface{} `json:"id" bson:"_id,omitempty"`
	Name      string      `json:"name" bson:"name"`
	Year      string      `json:"year" bson:"year"`
	Directors []string    `json:"directors" bson:"directors"`
	Writers   []string    `json:"writers" bson:"writers"`
	BoxOffice `json:"boxOffice" bson:"boxOffice"`
}

type BoxOffice struct {
	Budget uint64 `json:"budget" bson:"budget"`
	Gross  uint64 `json:"gross" bson:"gross"`
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(context.TODO())

	db := &DB{
		collection: client.Database("appDB").Collection("movies"),
	}

	r := mux.NewRouter()
	r.HandleFunc("/v1/movies/{id:[z-aA-Z0-0]*}", db.GetMovie).Methods(http.MethodGet)
	r.HandleFunc("/v1/movies", db.PostMovie).Methods(http.MethodPost)

	srv := &http.Server{
		Handler: r,
		Addr: ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
