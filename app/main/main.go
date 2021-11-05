package main

import (
	"github.com/gin-gonic/gin"
	"marbun.com/m/config"
	"marbun.com/m/controller"
	"net/http"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/books", controller.FindBooks)
	r.POST("/books", controller.CreateBook)
	r.GET("/books/:id", controller.FindBook)
	r.PATCH("/books/:id", controller.UpdateBook)
	r.DELETE("/books/:id", controller.DeleteBook)

	r.Run("0.0.0.0:8083")
}
