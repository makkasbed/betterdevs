package main

import (
	"github.com/gin-gonic/gin"
	"kasbedlabs.com/dbox-api/db"
	"kasbedlabs.com/dbox-api/routes"
)

func main() {
	//getting context
	defer db.Disconnect()
	router := gin.Default()
	router.POST("/users", routes.CreateUser)
	router.POST("/folders", routes.CreateFolder)
	router.GET("/folders/:id", routes.ListFolders)
	router.POST("/files", routes.CreateFile)
	router.POST("/login", routes.LoginUser)

	router.Run("localhost:18085")

}
