[![GoDoc](https://godoc.org/github.com/AquoDev/simple-imageboard-golang?status.svg)](https://godoc.org/github.com/AquoDev/simple-imageboard-golang)
[![Go Report Card](https://goreportcard.com/badge/github.com/AquoDev/simple-imageboard-golang)](https://goreportcard.com/report/github.com/AquoDev/simple-imageboard-golang)
[![Latest release](https://img.shields.io/github/v/release/AquoDev/simple-imageboard-golang)](https://github.com/AquoDev/simple-imageboard-golang/releases/latest)
![Code size in bytes](https://img.shields.io/github/languages/code-size/AquoDev/simple-imageboard-golang)
![License](https://img.shields.io/github/license/AquoDev/simple-imageboard-golang)

# Simple Imageboard

![Diagram](https://i.imgur.com/8YVWuRM.png)
**Everything is a post.**

# Table of contents

-   [Prerequisites](#prerequisites)
    -   [For container deployment](#for-container-deployment)
    -   [For local deployment](#for-local-deployment)
    -   [For both cases](#for-both-cases)
-   [Container deployment](#container-deployment)
    -   [First run: build containers and start containers in background](#first-run-build-containers-and-start-containers-in-background)
    -   [Start containers in foreground](#start-containers-in-foreground)
    -   [Start containers in background](#start-containers-in-background)
    -   [Stop containers in background](#stop-containers-in-background)
    -   [Stop and/or remove containers](#stop-andor-remove-containers)
    -   [Delete all saved data and remove containers](#delete-all-saved-data-and-remove-containers)
    -   [Rebuild containers from scratch](#rebuild-containers-from-scratch)
-   [Local deployment](#local-deployment)
    -   [Redis](#redis)
        -   [Set the same password in .env and redis.conf](#set-the-same-password-in-env-and-redisconf)
    -   [Database](#database)
        -   [Create database and user](#create-database-and-user)
        -   [Run migrations](#run-migrations)
    -   [Server](#server)
        -   [Option 1 (recommended): build and start server](#option-1-recommended-build-and-start-server)
        -   [Option 2: start server without building it](#option-2-start-server-without-building-it)
-   [Mixed deployment](#mixed-deployment)
    -   [Only Redis and Postgres as containers](#only-redis-and-postgres-as-containers)
-   [Tips](#tips)

# Prerequisites

### For container deployment

-   Docker
-   Docker Compose

### For local deployment

-   PostgreSQL
-   Redis
-   Golang (v1.13+)

### For both cases

You are required to run this command:

```console
cp .env.example .env
```

# Container deployment

#### First run: build containers and start containers in background

```console
docker-compose up --build -d
```

#### Start containers in foreground

```console
docker-compose up
```

#### Start containers in background

```console
docker-compose up -d
```

#### Stop containers in background

```console
docker-compose stop
```

#### Stop and/or remove containers

```console
docker-compose down
```

#### Delete all saved data and remove containers

```console
docker-compose down -v
```

#### Rebuild containers from scratch

```console
docker-compose build --no-cache
```

# Local deployment

### Redis

#### Set the same password in `.env` and `redis.conf`

```console
nano .env
...
REDIS_PASSWORD=your_pass_here
... (Save)
```

```console
sudo nano /etc/redis/redis.conf
...
# Uncomment requirepass
requirepass your_pass_here
... (Save)
```

### Database

#### Create database and user

```console
sudo -u postgres psql
> CREATE DATABASE simple_imageboard;
> CREATE USER username WITH ENCRYPTED PASSWORD 'password';
> GRANT ALL PRIVILEGES ON DATABASE simple_imageboard TO username;
> \q
```

#### Run migrations

`Migrations are automatically run on first start.`

### Server

#### Option 1 (recommended): build and start server

```console
go build
```

```console
./simple-imageboard-golang
```

#### Option 2: start server without building it

```console
go run main.go
```

# Mixed deployment

### Only Redis and Postgres as containers

```console
docker-compose -f docker-compose.yml -f docker-compose.mixed-deployment.yml up -d redis database
```

# Tips

-   First of all, **read `.env` and change the settings as you need**.
-   Use `autocannon` for performance testing.
    1. **Install:** `npm install -g autocannon`
    2. **Run:** `autocannon localhost:3000`
