package dbcon

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func ConnectDB() *mongo.Collection {

	
<<<<<<< HEAD
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
=======
	//clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	clientOptions := options.Client().ApplyURI("mongodb+srv://bookxadmin:bookx@bookxcluster-bcumy.mongodb.net/test?retryWrites=true&w=majority")
>>>>>>> new_branch
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

<<<<<<< HEAD
	collection := client.Database("book_information").Collection("books")
=======
	collection := client.Database("bookxcluster").Collection("books")
>>>>>>> new_branch

	return collection
}


type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}


func GetError(err error, w http.ResponseWriter) {

	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}