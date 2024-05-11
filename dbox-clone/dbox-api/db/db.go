package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"

	"kasbedlabs.com/dbox-api/models"
)

var client *mongo.Client
var collection *mongo.Collection
var COLLECTION = ""

func GetClient() *mongo.Client {
	uri := os.Getenv("DB_URL")
	//getting context
	if client != nil {
		return client
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//getting client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

func GetCollection(client *mongo.Client, collectioName string) *mongo.Collection {
	if collection != nil {
		return collection
	}
	collection := client.Database(os.Getenv("DB")).Collection(collectioName)
	return collection
}

func Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if client == nil {
		return
	}
	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}

func CreateUser(user models.User) string {
	client := GetClient()
	usersCollection := GetCollection(client, "users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	userObject := models.User{
		Id:       primitive.NewObjectID(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	result, err := usersCollection.InsertOne(ctx, userObject)
	if err != nil {
		return ""
	}
	return result.InsertedID.(primitive.ObjectID).Hex()
}
func Login(email string, password string) *models.User {
	client := GetClient()
	collection := GetCollection(client, "users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user *models.User
	filter := bson.D{{Name: "email", Value: email}}
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil
	}
	return user
}

func SaveFolder(folder models.Folder) string {
	client := GetClient()
	foldersCollection := GetCollection(client, "folders")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	folderObject := models.Folder{
		Id:        primitive.NewObjectID(),
		Name:      folder.Name,
		CreatedAt: time.Now().Format("2024-05-11 17:45"),
	}
	result, err := foldersCollection.InsertOne(ctx, folderObject)
	if err != nil {
		return ""
	}
	return result.InsertedID.(primitive.ObjectID).Hex()
}

func ListFolders(userId string) []models.Folder {
	client := GetClient()
	collection := GetCollection(client, "folders")
	//mongo queries
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var folders []models.Folder
	cursor, err := collection.Find(ctx, bson.D{{Name: "user_id", Value: userId}})
	defer cursor.Close(ctx)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	//Iterating through the book elements
	for cursor.Next(ctx) {
		var folder models.Folder
		err := cursor.Decode(&folder)
		if err != nil {
			log.Fatalln(err)
		}
		folder.Files = ListFiles(folder.Id.Hex())
		folders = append(folders, folder)
	}

	return folders
}

func SaveFile(file models.File) string {
	client := GetClient()
	filesCollection := GetCollection(client, "files")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fileObject := models.File{
		Id:        primitive.NewObjectID(),
		Name:      file.Name,
		FileType:  file.FileType,
		CreatedAt: time.Now().Format("2024-05-11 17:43"),
	}
	result, err := filesCollection.InsertOne(ctx, fileObject)
	if err != nil {
		return ""
	}
	return result.InsertedID.(primitive.ObjectID).Hex()
}

func ListFiles(folderId string) []models.File {
	client := GetClient()
	collection := GetCollection(client, "files")
	//mongo queries
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var files []models.File
	cursor, err := collection.Find(ctx, bson.D{{Name: "folder", Value: folderId}})
	defer cursor.Close(ctx)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	//Iterating through the book elements
	for cursor.Next(ctx) {
		var file models.File
		err := cursor.Decode(&file)
		if err != nil {
			log.Fatalln(err)
		}
		files = append(files, file)
	}

	return files
}
