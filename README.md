# Meme Coin API

## Table of Contents

- [Project Overview](#project-overview)
- [Project Structure](#project-structure)
- [Tech Stack](#tech-stack)
- [Prerequisites](#prerequisites)
- [Environment and Configuration](#environment-and-configuration)
- [Running the Application](#running-the-application)
    - [Docker (Recommended)](#docker-recommended)
    - [Locally](#locally)
- [API Overview](#api-overview)
- [Roadmap](#roadmap)

## Project Overview

This project implements a Meme Coin API using Golang with the Gin Web Framework and GORM. It provides functionalities
for creating, retrieving, updating, deleting, and "poking" meme coins (increasing their popularity). PostgreSQL is used
as the database, and the application is containerized using Docker and managed with docker-compose.

## Project Structure

```plaintext
.
├── cmd
│   └── server
│       └── main.go
├── config
│   └── config.go
├── docker-compose.yaml
├── Dockerfile
├── go.mod
├── go.sum
├── internal
│   ├── dtos
│   │   ├── meme_coin_request.go
│   │   └── meme_coin_response.go
│   ├── handlers
│   │   └── meme_coin_handler.go
│   ├── models
│   │   └── meme_coin.go
│   ├── repositories
│   │   └── meme_coin_repository.go
│   ├── routers
│   │   └── router.go
│   └── services
│       └── meme_coin_service.go
├── migrations
│   ├── 20250329134617_create_meme_coins_table.down.sql
│   └── 20250329134617_create_meme_coins_table.up.sql
└── pkg
    └── database
        └── db.go
```

- `cmd/server`: Application entry point, starts the HTTP server.
- `config`: Loads and manages environment variables and app configuration.
- `internal`: Core application logic (split into handlers, services, repositories, etc.)
  - `internal/dtos`: Defines request and response DTOs (Data Transfer Objects).
  - `internal/handlers`: Gin HTTP handlers (controller layer).
  - `internal/models`: GORM models representing DB schema.
  - `internal/repositories`: Contains logic for data persistence and database queries.
  - `internal/routers`: Initializes and groups application routes.
  - `internal/services`: Business logic layer, performs operations between handlers and repositories.
- `pkg/database`: Manages database connection setup.
- `migrations`: SQL migration scripts managed by `golang-migrate`.

## Tech Stack

- Go: [`1.24.1`](https://go.dev/doc/install) (managed by [gvm](https://github.com/moovweb/gvm))
- Web framework: [`Gin v1.10.0`](https://gin-gonic.com/)
- Database: [`PostgreSQL`](https://www.postgresql.org/)
- ORM: [`GORM`](https://gorm.io/)
- Migration CLI: [`golang-migrate`](https://github.com/golang-migrate/migrate)
- Containerization: [`Docker`](https://www.docker.com/)

## Prerequisites

- Go (`v1.24.1`)
- Docker (PostgreSQL will be launched via Docker Compose)
- [golang-migrate](https://github.com/golang-migrate/migrate)

## Environment and Configuration

Create a `.env` file (refer to `.env.example`) with these variables:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=meme_user
DB_PASSWORD=meme_pass
DB_NAME=meme_coin
SERVER_PORT=8080
API_VERSION=v1
```

## Running the Application

### Docker (Recommended)

1. Ensure your `.env` file is correctly set.
2. Run Docker Compose:

```bash
docker-compose up -d
```

3. Verify application is running via API tools such as Postman or cURL:

```
GET http://localhost:8080/health
```

### Locally

1. Install dependencies:

```bash
go mod download
```

2. Install Migration CLI:

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

3. Run PostgreSQL via Docker Compose:

```bash
docker-compose up -d postgres
```

4. Run database migrations:

```bash
./migrate.sh up 1
```

5. Run the application:

```bash
go run ./cmd/server/main.go
```

6. Verify local application running via API tools such as Postman or cURL:

```
GET http://localhost:8080/health
```

## API Overview

- **Create Meme Coin** (`POST /api/v1/meme-coins`):
    - Inputs: `name` (required), `description` (optional)
    - Creates a meme coin, initializes `popularity_score` to 0.

- **Get Meme Coin** (`GET /api/v1/meme-coins/:id`)

- **Update Meme Coin** (`PATCH /api/v1/meme-coins/:id`):
    - Inputs: `description`

- **Delete Meme Coin** (`DELETE /api/v1/meme-coins/:id`)

- **Poke Meme Coin** (`POST /api/v1/meme-coins/:id/poke`):
    - Increases the `popularity_score` by 1.

## Roadmap

- [ ] Implement Logger
- [ ] Write unit and integration tests

