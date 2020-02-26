# Start from Golang v1.13 base image
FROM golang:1.14.0-alpine3.11

# Download Git
RUN apk update && apk add --no-cache git

# Clone this repo
RUN git clone https://github.com/AquoDev/simple-imageboard-golang.git /app

# Change workdir
WORKDIR /app

# Copy .env.example to .env
RUN cp .env.example .env

# Build server
RUN go build -o server.bin

# Start server
CMD ["./server.bin"]
