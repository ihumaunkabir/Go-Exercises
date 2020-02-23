package main

import (

	"context"
	"encoding/json"
	"log"
	"net/http"
<<<<<<< HEAD
=======
	"os"
>>>>>>> new_branch

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/oasiscse/bookx/dbcon"
	"github.com/oasiscse/bookx/models"

)

func main() {

	r := mux.NewRouter()
<<<<<<< HEAD

  	
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{bookid}", searchByID).Methods("GET")
	r.HandleFunc("/api/booky/{year}", searchBook).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", r))
=======
	port := os.Getenv("PORT")

  	r.HandleFunc("/", homepage).Methods("GET")
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/bookid/{bookid}", searchByID).Methods("GET")
	r.HandleFunc("/api/bookdel/{bookid}", deleteByID).Methods("DELETE")
	r.HandleFunc("/api/booky/{year}", searchBook).Methods("GET")
	r.HandleFunc("/api/books/{title}/{author}/{psher}/{year}/{cat}/{bookid}", createURL).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+port, r))
>>>>>>> new_branch

}


<<<<<<< HEAD
=======
func homepage(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"bookx" : "A clever monkey is cooking for you, wait..."}`))
}


>>>>>>> new_branch
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


<<<<<<< HEAD
=======
func createURL(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	
	pms := mux.Vars(r)

	author := pms["author"]
	title := pms["title"]
	psher := pms["psher"]
	year := pms["year"]
	cat := pms["cat"]
	bookid := pms["bookid"]

	book := models.Book{
		Title : title,
		Publisher : psher,
		Author : author,
		Year : year,
		Category : cat,
		BookID : bookid,
	}	

	collection := dbcon.ConnectDB()
	result, err := collection.InsertOne(context.TODO(), book)

	if err != nil {
		w.Write([]byte(`{"message" : "Something went wrong."}`))
	} else{
		json.NewEncoder(w).Encode(result)
	}

}


>>>>>>> new_branch
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
<<<<<<< HEAD
		dbcon.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(book)
=======
		w.Write([]byte(`{"message": "Something went wrong."}`))
		//dbcon.GetError(err, w)
		//return

	} else{
		json.NewEncoder(w).Encode(book)
	}

>>>>>>> new_branch
}


func searchByID(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")

	var book models.Book
	var params = mux.Vars(r)
	bookid := string(params["bookid"])
	collection := dbcon.ConnectDB()

	filter := bson.M{"bookid": bookid}
	err := collection.FindOne(context.TODO(), filter).Decode(&book)

	if err != nil {
<<<<<<< HEAD
		dbcon.GetError(err, w)
		return

	}

	json.NewEncoder(w).Encode(book)


=======
		w.Write([]byte(`{"message": "Something went wrong."}`))
		//dbcon.GetError(err, w)
		//return

	} else{
		json.NewEncoder(w).Encode(book)
	}

>>>>>>> new_branch
}

func deleteByID(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)
<<<<<<< HEAD
	bookid := string(params["bookid"])
=======
	bookid := params["bookid"]
>>>>>>> new_branch
	collection := dbcon.ConnectDB()

	filter := bson.M{"bookid": bookid}
	delRes, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
<<<<<<< HEAD
		dbcon.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(delRes)

=======
		//dbcon.GetError(err, w)
		//return
		w.Write([]byte(`{"message": "Something went wrong."}`))
	} else{
		json.NewEncoder(w).Encode(delRes)
	}
>>>>>>> new_branch
}