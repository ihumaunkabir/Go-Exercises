package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type Book struct {
	ID     primitive.ObjectID 	`json:"_id,omitempty" bson:"_id,omitempty"`
    Title  string       `json:"title,omitempty" bson:"title,omitempty"`
	Author string       `json:"author, omitempty" bson:"author,omitempty"`
	Publisher   string  `json:"psher,omitempty" bson:"psher,omitempty"`
	Year string 		`json:"year,omitempty" bson:"year,omitempty"`
	Category string 	`json:"cat, omitempty" bson: "cat, omitempty"`
}
