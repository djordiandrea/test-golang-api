package main

import (
	"rest-api/controller"
	"rest-api/database"

	"github.com/gin-gonic/gin"
)

func main() {

	// dsn := "root:@tcp(127.0.0.1:3306)/warehouse?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	r := gin.Default()

	database.ConnectDatabase()

	v1 := r.Group("v1") //versioning, nanti r ini diganti sama v1

	r.GET("/user", controller.GetAllUser)
	r.GET("/user/:username", controller.GetUserDetail)

	r.GET("/", controller.RootHandler)
	v1.GET("/home", controller.RHome)
	r.GET("/books/:id", controller.RBook)
	// r.GET("/query", rQuery)
	r.GET("/books/:id/:title", controller.DoubleParam)
	r.GET("/query", controller.DoubleQuery)
	r.POST("/books", controller.PostBook)
	r.POST("/items", controller.PostWithValidation)

	r.Run(":8888")
}
