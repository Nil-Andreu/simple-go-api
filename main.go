package main  // This is the main module of the project

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// Define router GET method on the / path, as well as define a gin function for the JSON response
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"message" : "Hello world",
		})
	})

	// Deploy the router on the 8080 port
	router.Run()
}
