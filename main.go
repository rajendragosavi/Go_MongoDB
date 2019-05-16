package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
type Vote_Pole struct {
	Voter_Name string
	Party_Name string
}
*/
func main() {
	//set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("a", clientOptions)
	//Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//check the connection

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected Successfully to MongoDB")

}
