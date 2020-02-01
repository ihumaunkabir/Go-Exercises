package main

import (
    "context"
    "fmt"
    "log"

    //"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
    Name string
    Age  int
    City string
}

func main() {

    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil {
        log.Fatal(err)
    }


    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB!")


    collection := client.Database("test").Collection("info")

    ash := Trainer{"Ash", 10, "Pallet Town"}
    misty := Trainer{"Misty", 10, "Cerulean City"}
    brock := Trainer{"Brock", 15, "Pewter City"}
    

    insertResult, err := collection.InsertOne(context.TODO(), ash)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Inserted a single document: ", insertResult.InsertedID)


    trainers := []interface{}{misty, brock}
    insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
    
    if err != nil {
        log.Fatal(err)
    }


    fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)



    err = client.Disconnect(context.TODO())

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connection to MongoDB closed.")


}