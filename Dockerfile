# Start from Golang v1.11 base image
FROM golang:1.11-alpine

# Download Git
RUN apk add --no-cache git

# Download the dependencies
RUN go get -u -v github.com/AquoDev/simple-imageboard-golang/...

# Change workdir
WORKDIR /go/src/github.com/AquoDev/simple-imageboard-golang

# Rename .env file
RUN mv .env.example .env

# Run the server
CMD ["go", "run", "main.go"]
