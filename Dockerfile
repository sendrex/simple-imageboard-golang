##########################################################
# Start from Golang v1.14 base image to build the server #
##########################################################
FROM golang:1.14.6-alpine3.12 as build

# Download Git
RUN apk update && apk add --no-cache git

# Clone this repo
RUN git clone https://github.com/AquoDev/simple-imageboard-golang.git /app

# Change workdir
WORKDIR /app

# Checkout latest tag
RUN git checkout -q $(git tag --sort=taggerdate | tail -1)

# Build server
RUN CGO_ENABLED=0 go build -mod=vendor ./cmd/server-simple-imageboard

####################################
# Run the server in this container #
####################################
FROM alpine:3.12

WORKDIR /app

COPY --from=build /app/.env.example .env
COPY --from=build /app/static/ static/
COPY --from=build /app/server-simple-imageboard .

CMD ["./server-simple-imageboard"]
