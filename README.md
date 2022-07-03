# Simple API with Go Gin
Defining a simple api with Go Gin:
```
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
```
Gin provides with the router, where we can define the different endpoints and their responses.
If we wanted to run on a different port, we would put *r.Run(":5000")*.

To run this endpoint:
```
    go run main.go
```

And to test this endpoint:
```
    curl -X GET http://localhost:8080
```
