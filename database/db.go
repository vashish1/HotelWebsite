package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

//Data saves the data of the message
type Data struct {
	Fname string 
	Lname string 
	Email string 
	Msg   string 
}

//Createdb creates a database
func Createdb() (*mongo.Collection, *mongo.Client) {
	// Rest of the code will go here
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection := client.Database("HotelWebsite").Collection("Data")
	return collection, client
}

//Insertintodb inserts the data to the database
func Insertintodb(clctn *mongo.Collection, db *Data) {

	insertResult, err := clctn.InsertOne(context.TODO(), db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

}

//Findfromdb finds the required data
func Findfromdb(collection *mongo.Collection, st string) {
	filter := bson.D{{"Fname", "yashi"}}
	var result Data

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

//Disconnectdb disconnects the database
func Disconnectdb(client *mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
