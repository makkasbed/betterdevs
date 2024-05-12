package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"

	"kasbedlabs.com/dbox-api/models"
	"kasbedlabs.com/dbox-api/utils"
)

var client *mongo.Client
var collection *mongo.Collection

const DATE_FORMAT = "2006-01-02 15:04:05"

func GetClient() *mongo.Client {
	godotenv.Load()
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
	id := primitive.NewObjectID()
	hash, _ := utils.HashPassword(user.Password)
	userObject := models.User{
		Id:        id,
		Name:      user.Name,
		Email:     user.Email,
		Password:  hash,
		CreatedAt: time.Now().Format(DATE_FORMAT),
	}
	result, err := usersCollection.InsertOne(ctx, userObject)
	if err != nil {
		return ""
	}
	fmt.Println(result)
	return id.Hex()
}
func FindUser(id string) *models.User {
	client := GetClient()
	collection := GetCollection(client, "users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user *models.User
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"id": objID}
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil
	}
	return user
}
func Login(email string, password string) *models.User {
	client := GetClient()
	collection := GetCollection(client, "users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user *models.User
	filter := bson.M{"email": email}
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
	id := primitive.NewObjectID()
	folderObject := models.Folder{
		Id:        id,
		Name:      folder.Name,
		CreatedAt: time.Now().Format(DATE_FORMAT),
		UserId:    folder.UserId,
	}
	result, err := foldersCollection.InsertOne(ctx, folderObject)
	if err != nil {
		return ""
	}
	fmt.Println(result)
	return id.Hex()
}

func ListFolders(userId string) []models.Folder {
	client := GetClient()
	collection := GetCollection(client, "folders")
	//mongo queries
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var folders []models.Folder
	cursor, err := collection.Find(ctx, bson.M{"userid": userId})
	if err != nil {
		fmt.Println(err.Error())
	}
	defer cursor.Close(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	//Iterating through the book elements
	for cursor.Next(ctx) {
		var folder models.Folder
		err := cursor.Decode(&folder)
		if err != nil {
			fmt.Println(err.Error())
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
	id := primitive.NewObjectID()
	fileObject := models.File{
		Id:        id,
		Name:      file.Name,
		FileType:  file.FileType,
		CreatedAt: time.Now().Format(DATE_FORMAT),
		Folder:    file.Folder,
		UserId:    file.UserId,
	}
	result, err := filesCollection.InsertOne(ctx, fileObject)
	if err != nil {
		return ""
	}
	fmt.Println(result)
	return id.Hex()
}

func ListFiles(folderId string) []models.File {
	client := GetClient()
	collection := GetCollection(client, "files")
	//mongo queries
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var files []models.File
	cursor, err := collection.Find(ctx, bson.M{"folder": folderId})
	if err != nil {
		fmt.Println(err.Error())
	}
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
