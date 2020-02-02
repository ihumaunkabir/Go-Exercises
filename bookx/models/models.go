package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type Book struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Publisher   string        `json:"psher,omitempty" bson:"psher,omitempty"`
	Year string 			  `json:"year,omitempty" bson:"year,omitempty"`
	Title  string             `json:"title" bson:"title,omitempty"`
	Author *Author            `json:"author" bson:"author,omitempty"`
}

type Author struct {
	FirstName string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty" bson:"lastname,omitempty"`
}