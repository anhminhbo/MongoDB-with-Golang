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

var (
	collection *mongo.Collection
	ctx        = context.TODO()
)

func main() {
// insert
	minh := &Person{"Minh", 24}
	cici := &Person{"Cici", 25}
	junior := &Person{"Junior", 4}
	// insert one
	addMinh := addPerson(ctx, minh)
	CheckError(addMinh)

	//insert many
	persons := []interface{}{cici, junior}
	addMany := addManyPeople(ctx, persons)
	CheckError(addMany)

// update
	filter := bson.D{{"name", "Minh"}}

	update := bson.D{
		{"$set", bson.D{
			{"age", 30},
		}},
	}

	updateMinh := updatePerson(ctx, filter, update)
	CheckError(updateMinh)

// find a data
	var res Person
	respond, _ := findaPerson(ctx, filter, res)
	fmt.Println(respond)

// delete All data in the Database
	// deleteAll := deleteMany(ctx)
	// CheckError(deleteAll)
}

// Initialize mongo DB connection and Database
func init() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	CheckError(err)

	// Check the connection
	err = client.Ping(ctx, nil)
	CheckError(err)

	// get collection as ref
	collection = client.Database("testdb").Collection("people")
}
// Function to add one person
func addPerson(ctx context.Context, person *Person) error {
	_, err := collection.InsertOne(ctx, person)
	return err
}
//Function to add many people
func addManyPeople(ctx context.Context, persons []interface{}) error {
	_, err := collection.InsertMany(ctx, persons)
	return err
}
// Function to update a person 
func updatePerson(ctx context.Context, filter bson.D, update bson.D) error {
	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}
// Function to find a person
func findaPerson(ctx context.Context, filter bson.D, res Person) (Person, error) {
	e := collection.FindOne(ctx, filter).Decode(&res)
	return res, e
}
// Funtion to delete all data in the DB
// func deleteMany(ctx context.Context) error {
// 	_, err := collection.DeleteMany(ctx, bson.D{{}})
// 	return err
// }

// Function to check error
func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
