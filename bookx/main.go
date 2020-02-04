package main

import (

	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/oasiscse/bookx/dbcon"
	"github.com/oasiscse/bookx/models"

)

func main() {

	r := mux.NewRouter()

  	
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{year}", searchBook).Methods("GET")
  	
	log.Fatal(http.ListenAndServe(":3000", r))

}


func createBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var book models.Book

	_ = json.NewDecoder(r.Body).Decode(&book)

	collection := dbcon.ConnectDB()

	result, err := collection.InsertOne(context.TODO(), book)

	if err != nil {
		dbcon.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}


func getBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var books []models.Book

	collection := dbcon.ConnectDB()

	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		dbcon.GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var book models.Book

		err := cur.Decode(&book) 
		if err != nil {
			log.Fatal(err)
		}

		books = append(books, book)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(books) 
}


func searchBook(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")

	var book models.Book
	
	var params = mux.Vars(r)

	
	year := string(params["year"])
	collection := dbcon.ConnectDB()


	filter := bson.M{"year": year}
	err := collection.FindOne(context.TODO(), filter).Decode(&book)

	if err != nil {
		dbcon.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(book)
}
