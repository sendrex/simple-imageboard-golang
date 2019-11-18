# Simple Imageboard

No captcha. No botnet. No frontend.

## Prerequisites

### For local deployment

-   PostgreSQL
-   Redis
-   Golang (v1.13+)

Run these commands:
```console
cp .env.example .env
go get -u -v ./...
```


### For container deployment

-   Docker
-   Docker Compose

---

## Container deployment

### First run: build containers and start containers in background

```console
docker-compose up --build -d
```

### Start containers in foreground

```console
docker-compose up
```

### Start containers in background

```console
docker-compose up -d
```

### Stop containers in background

```console
docker-compose stop
```

### Stop and/or remove containers

```console
docker-compose down
```

### Delete all saved data and remove containers

```console
docker-compose down -v
```

### Rebuild containers from scratch

```console
docker-compose build --no-cache
```

---

## Local deployment

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

#### Start server

```console
go run main.go
```

---

## Tips

-   First of all, **read `.env` and change the settings as you need**.
-   You are supposed to know what you are doing, right?
-   Use `autocannon` for performance testing.
    1. **Install:** `npm install -g autocannon`
    2. **Run:** `autocannon http://localhost:3000`
