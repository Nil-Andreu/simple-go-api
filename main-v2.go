package main

import (
	"github.com/gin-gonic/gin"
)

func SimpleRequestHandler(c *gin.Context) {
	c.JSON(200, gin.H {
		"message" : "Hello world!",
	})
}

func NameRequestHandler(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H {
		"message" : "Hello " + name,
	})
}

func main() {
	router := gin.Default()
	router.GET("/", SimpleRequestHandler)
	router.GET("/name/:name", NameRequestHandler)
	router.Run(":8080")
}