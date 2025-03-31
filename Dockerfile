# Build stage
FROM golang:1.24.1-alpine AS builder
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
# This layer is cached if the dependency files don't change
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code into the container
COPY . .

# Compile the application
# Using CGO_ENABLED=0 for a static binary that works in the minimal alpine image
RUN CGO_ENABLED=0 GOOS=linux go build -o meme-coin ./cmd/server

# Run stage - using a minimal alpine image for smaller final image size
FROM alpine:latest
WORKDIR /root/

# Install curl and unzip, then download and configure migrate CLI
# Combined into a single RUN to reduce image layers
RUN apk add --no-cache curl unzip && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.2/migrate.linux-amd64.tar.gz | tar xvz -C /tmp && \
    mv /tmp/migrate /migrate && \
    chmod +x /migrate

# Copy the compiled binary, migration files, and startup scripts from the builder stage
COPY --from=builder /app/meme-coin .
COPY migrations /migrations
COPY wait-for-it.sh entrypoint.sh /
RUN chmod +x /wait-for-it.sh /entrypoint.sh

# Expose the application port
EXPOSE 8080

# Define the entrypoint script that will run when the container starts
ENTRYPOINT ["/entrypoint.sh"]