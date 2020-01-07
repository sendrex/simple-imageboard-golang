# Fourth version: update 1

### Added

-   `CHANGELOG.md`
-   `/health` and `/settings` endpoints.

### Removed

-   `static/info.txt` (replaced by `/settings`).

### Refactored

-   `model.Post` validation.
-   `redis/json.go` functions into `redis.setCachedModel` and `redis.getCachedModel`.

# Fourth version

### Updated

-   `README` with new requeriments
-   Cleaner queries in `database/post.go`
-   Renamed containers

# Third version: update 3

### Updated

-   Containers are unprivileged
-   Use of `model.Cache` in `handler.DeletePost()`

# Third version: update 2

### Updated

-   `api-si-go` container builds and runs the server as a binary (lower memory footprint)

### Added

-   Mixed deployment use case

# Third version: update 1

### Fixed

-   Some response errors in `handler.DeletePost()`

# Third version

### Added

-   `Cache` model to cache errors the same way as pages/posts/threads

### Removed for good

-   Custom error handling (every error is now handled by Echo)
-   `status` field from error responses

### Refactored

-   `godotenv` is only imported in `database/client.go` (it's the first init)
-   `redis` package uses `model.Cache` to unmarshal/marshal data
-   `util` package was removed, but `util.RandomString` was refactored into `model.Post.GenerateDeleteCode()`
-   Iris was replaced by Echo, meaning that `handler` and `middleware` packages had to be refactored
-   Better HTTP codes (as in more precise)

### Updated

-   `go.mod` and `go.sum`

# Second version

### Updated

-   `go.mod` and `go.sum`

### Refactored

-   `secure` middleware from `main.go` to `middleware` package

### Implemented missing features from v1.5

-   CORS
-   More `.env` variables (such as number of posts required to reach bump limit)

# First version: update 5

### Added

-   A diagram to the `README` that explains the structure of the imageboard

### Refactored

-   `redis` package now uses generic interfaces and maps to marshal/unmarshal data
-   `Post` model validation data outside the handler

### Implemented missing features from v1.4

-   Versioning for dependencies

# First version: update 4

### Added

-   `replies` field to model `Post` (it's the only field that's not saved in the table)

### Refactored

-   GET queries in `database` package to count `replies` in every post

### Implemented missing features from v1.3

-   Nothing, but some important fixes and the addition of `replies` are worth a new release

# First version: update 3

### Refactored

-   All packages have a depth level of 1
-   `redis` and `database` are the most modular packages (in fact, their clients are not used outside those packages)
-   `redis` has been split into several files, it provides better cache management
-   `model` contains 2 abstracted structs: `Post` and `DeleteData`, so they can be used across the project
-   GET methods from `handler` follow the "cache-database-error" path more clearly

### Implemented missing features from v1.2

-   Nothing, but the refactoring is worth a new release

# First version: update 2

### Added

-   Secure headers for every request

### Implemented missing features from v1.1

-   Limiters

# First version: update 1

### Added

-   `LICENSE` (GPL-3.0)

### Refactored

-   Connection clients (Redis and GORM)
-   Go structs from `database/methods` and `server` (such as "Thread" or "Page") were removed
-   `ctx.WriteString` was changed to `ctx.JSON`
-   `middleware.UseContentTypeJSON` was redundant, so it was removed

### Implemented missing features from v1.0

-   `robots.txt`
-   A route where some information about the instance can be obtained (almost everything defined in `.env`)

# First version

### Added

-   Basic functionality
