# Start from Golang v1.13 base image
FROM golang:1.13.4-alpine3.10

# Download Git
RUN apk update && apk add --no-cache git

# Clone this repo
RUN git clone https://github.com/AquoDev/simple-imageboard-golang.git /simple-imageboard-golang

# Change workdir
WORKDIR /simple-imageboard-golang

# Make .env file from .env.example
RUN cp .env.example .env

# Build server
RUN go build

# Start server
CMD ["./simple-imageboard-golang"]
