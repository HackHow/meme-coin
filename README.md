# Meme Coin API

This is a Meme Coin API project developed using Golang and the Gin framework. This project provides a RESTful API for managing meme coins with a clear separation of concerns (handlers, models, services, repositories, middleware).

## Features

- Create, read, update, and delete meme coins.
- Poke a meme coin to increase its popularity score.
- Unified error handling and API response format.
- PostgreSQL as the database with GORM as the ORM.
- Docker containerization support (to be added in later stages).

## Requirements

- Golang 1.24.1
- PostgreSQL (can be run via Docker)
- Docker (for containerization)

## Getting Started

1. Clone the repository.
2. Navigate to the project directory.
3. Run `go run ./cmd/server` to start the server.

## Go Version Notes

If you don't have Go 1.24.1 installed, please download it from [Go Downloads](https://golang.org/dl/). If you're using gvm, be aware that you might encounter installation order issues (e.g., you may need to install an older version first).

