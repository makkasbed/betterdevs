package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"kasbedlabs.com/dbox-api/db"
	"kasbedlabs.com/dbox-api/models"
	"kasbedlabs.com/dbox-api/storage"
)

func CreateUser(c *gin.Context) {
	var user models.User
	var response models.Response
	err := c.Bind(&user)

	if err != nil {

		response.Status = 0
		response.Message = err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	id := db.CreateUser(user)

	if id != "" {
		isCreated := storage.MakeDirectory(id)
		if isCreated {
			fmt.Println("Folder created!")
		}
		response.Status = 1
		response.Message = "User created successfully!"
		response.Id = id
		c.JSON(http.StatusCreated, response)
	} else {
		response.Status = 0
		response.Message = "User creation failed!"
		c.JSON(http.StatusBadGateway, response)
	}
}

func CreateFolder(c *gin.Context) {
	var folder models.Folder
	var response models.Response
	err := c.Bind(&folder)

	if err != nil {

		response.Status = 0
		response.Message = err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	id := db.SaveFolder(folder)

	if id != "" {

		status := storage.MakeDirectory(folder.UserId + "/" + id)
		if status {
			fmt.Println("Folder created!")
		}
		response.Status = 1
		response.Message = "Folder created successfully!"
		response.Id = id
		c.JSON(http.StatusCreated, response)
	} else {
		response.Status = 0
		response.Message = "Folder creation failed!"
		c.JSON(http.StatusBadGateway, response)
	}
}

func CreateFile(c *gin.Context) {
	var file models.File
	var response models.Response
	err := c.Bind(&file)

	if err != nil {

		response.Status = 0
		response.Message = err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	id := db.SaveFile(file)

	if id != "" {
		path, err := storage.CreateFile(os.Getenv("STORE")+"/"+file.UserId+"/"+file.Folder, file.Name, file.FileType, file.Data)
		if err != nil {
			fmt.Println(err.Error())
		}
		response.Status = 1
		response.Message = "File created successfully: " + path
		response.Id = id
		c.JSON(http.StatusCreated, response)
	} else {
		response.Status = 0
		response.Message = "File creation failed!"
		c.JSON(http.StatusBadGateway, response)
	}
}

func ListFolders(c *gin.Context) {

	var response models.Response
	userID := c.Param("id")
	fmt.Println(userID)

	folders := db.ListFolders(userID)

	if len(folders) > 0 {
		response.Status = 1
		response.Message = "Folders available!"
		response.Folders = folders
		c.JSON(http.StatusCreated, response)
	} else {
		response.Status = 0
		response.Message = "Folder listing failed!"
		c.JSON(http.StatusBadGateway, response)
	}
}

func LoginUser(c *gin.Context) {
	var login models.Login
	var response models.Response

	err := c.Bind(&login)

	if err != nil {
		response.Status = 0
		response.Message = err.Error()

		c.JSON(http.StatusBadRequest, response)
	} else {
		user := db.Login(login.Email, login.Password)
		if user != nil {
			response.Status = 1
			response.Message = "Login successful!"
			c.JSON(http.StatusOK, response)
		} else {
			response.Status = 0
			response.Message = "Login unsuccessful!"
			c.JSON(http.StatusBadGateway, response)
		}

	}
}
