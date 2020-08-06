# Simple Imageboard

[![GoDev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go)](https://pkg.go.dev/github.com/AquoDev/simple-imageboard-golang?tab=overview)
[![Go report card](https://goreportcard.com/badge/github.com/AquoDev/simple-imageboard-golang)](https://goreportcard.com/report/github.com/AquoDev/simple-imageboard-golang)
[![Latest release](https://img.shields.io/github/v/release/AquoDev/simple-imageboard-golang?logo=github)](https://github.com/AquoDev/simple-imageboard-golang/releases/latest)
![License](https://img.shields.io/github/license/AquoDev/simple-imageboard-golang)

![Diagram](https://i.imgur.com/MsP4QU4.png)
**Everything is a post.**

## Flow diagram 🔄

![Diagram](https://i.imgur.com/90MK3y8.png)

## Prerequisites 📋

- Git
- Docker (19.03+)
- Docker Compose (1.25+)

**For local or mixed deployment**, you also need the following:

- PostgreSQL (12+)
- Redis (6.0+)
- Golang (1.14.4+)

## Installation 🔧

Open a terminal and follow these steps:

```console
# Clone repository
user@system:~$ git clone https://github.com/AquoDev/simple-imageboard-golang.git

# Change directory
user@system:~$ cd simple-imageboard-golang

# Copy .env.example and rename it to .env
user@system:simple-imageboard-golang$ cp .env.example .env

# Edit .env and change every value you need, in your editor of choice
user@system:simple-imageboard-golang$ editor .env
```

## Deployment 🚀

`It is assumed that you're in the same directory where the repository was cloned to.`

### Docker

```console
# Start containers
docker-compose up -d
```

```console
# Stop containers
docker-compose stop
```

```console
# Update server container without losing data
docker-compose build --no-cache
```

```console
# Remove containers without losing data
docker-compose down
```

```console
# Delete all saved data and remove containers
docker-compose down -v
```

### Local

#### Redis: set password

The credentials must be shared between `.env` and `/etc/redis/redis.conf`.

```console
editor .env
...
REDIS_PASSWORD=your_pass_here
... (Save)

sudo editor /etc/redis/redis.conf
...
# Uncomment requirepass
requirepass your_pass_here
... (Save)
```

#### Postgres: create user and database

The credentials must be shared between `.env` and these commands.

Tables are automatically created after starting the server for the first time.

```console
sudo -u postgres psql
> CREATE DATABASE simple_imageboard;
> CREATE USER username WITH ENCRYPTED PASSWORD 'password';
> GRANT ALL PRIVILEGES ON DATABASE simple_imageboard TO username;
> \q
```

#### Server: build and run

You can edit the listening port in `.env` and put a reverse proxy in front of this server.

Dependencies are bundled with the project (`vendor` directory), but if you wish to download them, use the online method.

```console
# Method 1: Using the bundled dependencies
go build -mod=vendor ./cmd/server-simple-imageboard

# Method 2: Download them (it can take a while)
go build ./cmd/server-simple-imageboard
```

```console
# Start server
./server-simple-imageboard
```

### Mixed

```console
# Start Redis container
docker-compose -f docker-compose.yml -f docker-compose.mixed-deployment.yml up -d redis
```

```console
# Start Postgres container
docker-compose -f docker-compose.yml -f docker-compose.mixed-deployment.yml up -d database
```

```console
# Start Redis and Postgres containers
docker-compose -f docker-compose.yml -f docker-compose.mixed-deployment.yml up -d redis database
```

## License 📋

[GNU General Public License v3.0](https://github.com/AquoDev/simple-imageboard-golang/blob/master/LICENSE)
