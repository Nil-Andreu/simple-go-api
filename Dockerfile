FROM golang:1.18-alpine

# Create the application default directory
WORKDIR /app

# Pass to this new directory the files we need
COPY main-v2.go .
COPY go.mod .
COPY go.sum .

# Download the dependencies
RUN go mod download

# Run the application
EXPOSE 8080
CMD ["go", "run", "main-v2.go"]
