# Start from Golang v1.13 base image
# NOTE: It can't be updated to Golang 1.14 until the first issue is fixed:
#   - https://github.com/golang/go/issues/37436
#   - https://github.com/docker-library/golang/issues/320
FROM golang:1.13.8-alpine3.11

# Download Git
RUN apk update && apk add --no-cache git

# Clone this repo
RUN git clone https://github.com/AquoDev/simple-imageboard-golang.git /app

# Change workdir
WORKDIR /app

# Checkout latest tag
RUN git checkout -q $(git tag --sort=taggerdate | tail -1)

# Copy .env.example to .env
RUN cp .env.example .env

# Build server
RUN go build -o server.bin

# Start server
CMD ["./server.bin"]
