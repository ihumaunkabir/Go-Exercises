package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var templ = template.Must(template.ParseFiles("index.html"))

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods("GET")

	fmt.Println("Running on 3000")
	log.Fatal(http.ListenAndServe(":3000", r))

}

func index(w http.ResponseWriter, r *http.Request) {
	templ.Execute(w, nil)
}
