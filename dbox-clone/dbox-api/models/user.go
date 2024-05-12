package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
	CreatedAt string             `json:"created_at"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
