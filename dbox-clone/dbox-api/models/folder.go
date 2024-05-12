package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Folder struct {
	Id        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	UserId    string             `json:"user_id"`
	CreatedAt string             `json:"created_at"`
	Files     []File
}

type File struct {
	Id        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	FileType  string             `json:"file_type"`
	Folder    string             `json:"folder"`
	CreatedAt string             `json:"created_at"`
	Data      string             `json:"data"`
	UserId    string             `json:"user_id"`
}
