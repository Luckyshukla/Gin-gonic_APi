package main

import (
	"github.com/gin-gonic/gin"
	"main.go/Controller"
	"main.go/Models"

)

func main() {
	r := gin.Default()
	
	
	
	r.POST("/login", Controller.Login)
	
	
	Models.ConnectDataBase()

	crud:= r.Group("/crud")
	r.Use()
	{
	crud.GET("/books", Controller.FindBooks)
	crud.POST("/books", Controller.CreateBook)
	crud.POST("/books/:id", Controller.FindBook)
	crud.PATCH("/books/:id", Controller.UpdateBook)
	crud.DELETE("/books/:id", Controller.DeleteBook)
	}
	r.Run()
}
