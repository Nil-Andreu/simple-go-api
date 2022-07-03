# Simple API with Go Gin
Defining a simple api with Go Gin:
```{go}
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

## Define the Go Modules

To define the dependencies/modules of our project, we put the following:
```{go}
    go mod init <name>
```

Then we go to the new go.mod created file, and put the dependencies.
```{go}
    module github.com/Nil-Andreu/simple-go-api

    go 1.18

    require github.com/gin-gonic/gin v1.8.1
```
So we are basically putting which is our github repository (to download the code from), which golang version we are using, and then which packages we need (Gin only in this case).

This will make all the work or then handling sub-dependencies.

## Running the app with Docker

Once we defined the dependencies, we can create the Docker image that we could use for then deploy this simple endpoint to a server.

For this, we use the Dockerfile:
```{go}
    FROM golang:1.18-alpine

    # Create the application default directory
    WORKDIR /app

    # Pass to this new directory the files we need
    COPY main.go .
    COPY go.mod .
    COPY go.sum .

    
    # Download the dependencies
    RUN go mod download

    # Run the application
    EXPOSE 8080
    CMD go run main.go
```
Where first we state which is the golang version we want to use.

Then we define which is going to be the WORKDIR of our app, and we copy the files that we need.

Then, we run the *go mod download* for installing the necessary dependencies.

Finally, exposing the endpoint and running the file where we have defined our endpoint.

So if we want to then create the image, and start a container, we can run at the terminal the following commands:
```
    docker build . simple-api:v1
    docker run -p 8080:8080 simple-api:v1 -d
```
For simplicity, I have put those commands in a Makefile. So you can run them directly with: make docker.

So we could go to the terminal on port 8080, and obtain the response.

# Adding another Enpoint
To give a bit more information, we could add another endpoint.
This one, checks for parameters on the URL to then create a custom output:
```
    import ("fmt"
    ...

    router.GET("/name/:name", func(c *gin.Context) {
		name := c.Param("name")
		message := fmt.Sprintf("Hello %s!", name)
		c.JSON(200, gin.H{
			"message": message,
		})
	})
```

