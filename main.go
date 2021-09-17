package main

import (
	"github.com/gin-gonic/gin"
	"main.go/Controller"
	"main.go/Models"
	//"github.com/go-redis/redis/v7"
    //"github.com/twinj/uuid"
)

func main() {
	r := gin.Default()
	Models.ConnectDataBase()

	//Basic Auth
	/*
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin":"password",
	}))
	*/

	r.POST("/login", Controller.Login)
	
	r.GET("/books", Controller.FindBooks)
	r.POST("/books", Controller.CreateBook)
	r.POST("/books/:id", Controller.FindBook)
	r.PATCH("/books/:id", Controller.UpdateBook)
	r.DELETE("/books/:id", Controller.DeleteBook)
	
	r.Run()
}
