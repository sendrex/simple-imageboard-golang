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
-   [Installation](#installation)
    -   [Clone this repository](#clone-this-repository)
    -   [Copy `.env.example` to `.env`](#copy-envexample-to-env)
    -   [Edit `.env` with your credentials](#edit-env-with-your-credentials)
-   [Container deployment](#container-deployment)
    -   [First run: build containers and start containers in background](#first-run-build-containers-and-start-containers-in-background)
    -   [Start containers in foreground](#start-containers-in-foreground)
    -   [Start containers in background](#start-containers-in-background)
    -   [Stop containers in background](#stop-containers-in-background)
    -   [Update containers without losing data](#update-containers-without-losing-data)
    -   [Remove containers without losing data](#remove-containers-without-losing-data)
    -   [Delete all saved data and remove containers](#delete-all-saved-data-and-remove-containers)
-   [Local deployment](#local-deployment)
    -   [Redis: set and share password](#redis-set-and-share-password)
    -   [Database: create database and user](#database-create-database-and-user)
    -   [Server: build and start it](#server-build-and-start-it)
-   [Mixed deployment](#mixed-deployment)
    -   [Only Redis as container](#only-redis-as-container)
    -   [Only Postgres as container](#only-postgres-as-container)
    -   [Redis and Postgres as containers](#redis-and-postgres-as-containers)
-   [Tips](#tips)

# Prerequisites

### For container deployment

-   Docker (v19.03+)
-   Docker Compose (v1.25+)

### For local deployment

-   PostgreSQL (v12+)
-   Redis (v5.0+)
-   Golang (v1.13+, but **not** v1.14+)

# Installation

### Clone this repository

```console
git clone https://github.com/AquoDev/simple-imageboard-golang.git
```

### Copy `.env.example` to `.env`

```console
cp .env.example .env
```

### Edit `.env` with your credentials

```console
nano .env
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

#### Update containers without losing data

```console
docker-compose build --no-cache
```

#### Remove containers without losing data

```console
docker-compose down
```

#### Delete all saved data and remove containers

```console
docker-compose down -v
```

# Local deployment

### Redis: set and share password

The credentials must be shared between `.env` and `/etc/redis/redis.conf`.

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

### Database: create database and user

The credentials must be shared between `.env` and these commands.

```console
sudo -u postgres psql
> CREATE DATABASE simple_imageboard;
> CREATE USER username WITH ENCRYPTED PASSWORD 'password';
> GRANT ALL PRIVILEGES ON DATABASE simple_imageboard TO username;
> \q
```

`Tables are automatically created after starting the server for the first time.`

### Server: build and start it

You can edit the listening port in `.env` and put a reverse proxy in front of this server.

```console
go build -o server.bin
```

```console
./server.bin
```

# Mixed deployment

### Only Redis as container

```console
docker-compose -f docker-compose.yml -f docker-compose.mixed-deployment.yml up -d redis
```

### Only Postgres as container

```console
docker-compose -f docker-compose.yml -f docker-compose.mixed-deployment.yml up -d database
```

### Redis and Postgres as containers

```console
docker-compose -f docker-compose.yml -f docker-compose.mixed-deployment.yml up -d redis database
```

# Tips

-   First of all, **read `.env` and change the settings as you need**.
-   Use `autocannon` for performance testing.
    1. **Install:** `npm install -g autocannon`
    2. **Run:** `autocannon localhost:3000`
