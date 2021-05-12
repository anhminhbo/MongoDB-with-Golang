package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, e := mongo.Connect(context.TODO(), clientOptions)
	CheckError(e)

	// Check the connection
	e = client.Ping(context.TODO(), nil)
	CheckError(e)

	// get collection as ref
	collection := client.Database("testdb").Collection("people")

	// insert
	minh := Person{"Minh", 24}
	cici := Person{"Cici", 25}
	junior := Person{"Junior", 4}

	_, e = collection.InsertOne(context.TODO(), minh)
	CheckError(e)

	persons := []interface{}{cici, junior}
	_, e = collection.InsertMany(context.TODO(), persons)
	CheckError(e)

	// update
	filter := bson.D{{"name", "Minh"}}

	update := bson.D{
		{"$set", bson.D{
			{"age", 30},
		}},
	}

	_, e = collection.UpdateOne(context.TODO(), filter, update)
	CheckError(e)

	// find
	var res Person
	e = collection.FindOne(context.TODO(), filter).Decode(&res)
	fmt.Println(res)

	// delete
	_, e = collection.DeleteMany(context.TODO(), bson.D{{}})
	CheckError(e)
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
