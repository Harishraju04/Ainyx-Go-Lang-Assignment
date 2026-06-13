# Ainyx Golang API

Ainyx Golang API is a simple RESTful user management service built with Go, Fiber, PostgreSQL, and sqlc. It includes request validation, structured logging, middleware for request ID injection and duration logging, and paginated user listing.

## Features

- REST API for user CRUD operations
- PostgreSQL backend with sqlc-generated queries
- Request validation using `go-playground/validator/v10`
- Structured logging with `zap`
- Global middleware for:
  - `X-Request-Id` response header
  - request duration logging
- Paginated user listing with `page` and `pageSize`

## API Endpoints

| Method | Path | Description |
|---|---|---|
| GET | `/api/health` | Health check endpoint |
| POST | `/api/users` | Create a new user |
| GET | `/api/users` | List users with pagination |
| GET | `/api/users/:id` | Get a user by ID |
| PUT | `/api/users/:id` | Update a user |
| DELETE | `/api/users/:id` | Delete a user |

### Pagination for `/api/users`

Query parameters:
- `page` (integer, default `1`)
- `pageSize` (integer, default `10`, max `100`)

Example:
```bash
GET /api/users?page=2&pageSize=20
```

Response includes:
- `data` — array of users
- `page` — current page
- `pageSize` — number of users per page
- `total` — total user count
- `totalPages` — total pages available

## Folder Structure

```
cmd/server/main.go
config/config.go
/db
  inti.go
  migrations/
  queries/queries.sql
  sqlc/
internal/
  handler/
  logger/
  middleware/
  repository/
  routes/
  service/
  validator/
sqlc.yaml
```

### Key directories

- `cmd/server/` — application entry point
- `config/` — environment and configuration loader
- `db/` — database initialization, migrations, and sqlc query definitions
- `internal/handler/` — HTTP request handlers
- `internal/service/` — business logic and pagination
- `internal/repository/` — data access layer
- `internal/middleware/` — request ID + duration middleware
- `internal/logger/` — Zap logger initialization
- `internal/routes/` — route setup and middleware registration
- `internal/validator/` — validator initialization

## Environment Variables

Required:
- `DATABASE_URL` — PostgreSQL connection string

Optional:
- `PORT` — server port (default `:8080`)

Example:
```env
DATABASE_URL=postgres://ainyx:ainyxpass@localhost:5432/ainyx?sslmode=disable
PORT=:8080
```

## Local Setup

1. Clone the repository:
```bash
git clone <repo-url> ainyx-golang-api
cd ainyx-golang-api
```

2. Create a `.env` file or export variables:
```env
DATABASE_URL=postgres://ainyx:ainyxpass@localhost:5432/ainyx?sslmode=disable
PORT=:8080
```

3. Install dependencies and build:
```bash
go mod download
go run ./cmd/server
```

4. Open your browser or API client:
- `http://localhost:8080/api/health`
- `http://localhost:8080/api/users`

## Docker

Build and run the container:
```bash
docker build -t ainyx-api .
docker run -e DATABASE_URL="postgres://ainyx:ainyxpass@db:5432/ainyx?sslmode=disable" -e PORT=":8080" -p 8080:8080 ainyx-api
```

> Note: `DATABASE_URL` must point to a running PostgreSQL instance.

## Notes

- The API injects `X-Request-Id` into every response header.
- Request duration is logged for each request.
- Validation errors return structured HTTP 400 responses.
- The user list endpoint uses server-side pagination.

## License

This project is provided as-is. Update license details as needed.
