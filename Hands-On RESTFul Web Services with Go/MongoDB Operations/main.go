package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Movie struct {
	Name string `bson:"name"`
	Year string `bson:"year"`
	Directors []string `bson:"directors"`
	Writers []string `bson:"writers"`
	BoxOffice `bson:"boxOffice"`
}

type BoxOffice struct {
	Budget uint64 `bson:"budget"`
	Gross uint64 `bson:"gross"`
}

func createMovie(collection *mongo.Collection, movie Movie) error {
	_, err := collection.InsertOne(context.TODO(), movie)
	return err
}

func getMovie(collection *mongo.Collection, filter bson.M) (Movie, error) {
	result := collection.FindOne(context.TODO(), filter)
	movie := Movie{}
	err := result.Decode(&movie)

	return movie, err
}

func main() {
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connect to mongodb server successfully")
	collection := client.Database("appDB").Collection("movies")
	
	movie := Movie{
		Name:      "The Dark KNight",
		Year:      "2008",
		Directors: []string{"Christopher Nolan"},
		Writers:   []string{"Jonathan Nolan", "Christopher Nolan"},
		BoxOffice: BoxOffice{
			Budget: 185000000,
			Gross: 533316061,
		},
	}


	err = createMovie(collection, movie)
	if err != nil {
		log.Fatal(err)
	}

	queryMovie, err := getMovie(collection, bson.M{
		"boxOffice.budget": bson.M{
			"$gt": 1800,
		},
	})
	if err != nil {
		log.Fatal("getMovie - ", err)
	}

	fmt.Println("movie:", queryMovie)

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}





}
