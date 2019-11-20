# Start from Golang v1.13 base image
FROM golang:1.13.4-alpine3.10

# Download Git
RUN apk update && apk add --no-cache git

# Download the dependencies
RUN go get -v github.com/AquoDev/simple-imageboard-golang

# Change workdir
WORKDIR /go/src/github.com/AquoDev/simple-imageboard-golang

# Rename .env file
RUN cp .env.example .env

# Run the server
CMD ["go", "run", "main.go"]
