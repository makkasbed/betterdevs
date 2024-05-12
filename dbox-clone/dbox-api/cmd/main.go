package main

import (
	"github.com/gin-gonic/gin"
	//ginSwagger "github.com/swaggo/gin-swagger"
	//"github.com/swaggo/gin-swagger/swaggerFiles"
	"kasbedlabs.com/dbox-api/db"
	"kasbedlabs.com/dbox-api/middlewares"
	"kasbedlabs.com/dbox-api/routes"
)

func main() {
	//getting context
	defer db.Disconnect()
	router := gin.Default()
	router.POST("/users", routes.CreateUser)
	router.POST("/folders", middlewares.Authorize, routes.CreateFolder)
	router.GET("/folders/:id", middlewares.Authorize, routes.ListFolders)
	router.POST("/files", middlewares.Authorize, routes.CreateFile)
	router.POST("/login", routes.LoginUser)

	//url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run("localhost:18085")

}
