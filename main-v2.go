package main

import (
	"github.com/gin-gonic/gin"
)

func FunctionHandler(c *gin.Context) {
	c.JSON(200, gin.H {
		"message" : "message",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", FunctionHandler)
	router.Run(":8080")
}