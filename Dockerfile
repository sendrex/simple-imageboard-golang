####################################################
# Start from Golang base image to build the server #
####################################################
FROM golang:1.15.7-alpine as build

# Read the following link for more info: https://github.com/moby/moby/issues/34513#issuecomment-389467566
LABEL stage=intermediate

# Copy this repo
COPY . /app

# Change workdir
WORKDIR /app

# Build server
RUN go build -mod=vendor ./cmd/server-simple-imageboard

####################################
# Run the server in this container #
####################################
FROM alpine:3.13

WORKDIR /app

COPY --from=build /app/.env.example .env
COPY --from=build /app/static/ .
COPY --from=build /app/server-simple-imageboard .

CMD ["./server-simple-imageboard"]
