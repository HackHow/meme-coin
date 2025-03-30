#!/bin/sh
set -e

echo "Running database migrations..."

/wait-for-it.sh meme_coin_postgres:5432 --timeout=30 --strict -- echo "Database started successfully!"
/migrate -path=/migrations -database="${DATABASE_URL}" up

echo "Starting application..."
exec ./meme-coin
