# Build stage
FROM golang:1.24.1-alpine AS builder
WORKDIR /app

# 安裝 git（如果需要）
RUN apk update && apk add --no-cache git

# 複製 go.mod 與 go.sum 並下載依賴
COPY go.mod go.sum ./
RUN go mod download

# 複製所有程式碼到容器中
COPY . .

# 編譯應用程式，假設 main package 位於 cmd/server 目錄中
RUN CGO_ENABLED=0 GOOS=linux go build -o meme-coin ./cmd/server

# Run stage
FROM alpine:latest
WORKDIR /root/

# 安裝 curl 與 unzip 並下載 migrate CLI
RUN apk add --no-cache curl unzip && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.2/migrate.linux-amd64.tar.gz | tar xvz -C /tmp && \
    mv /tmp/migrate /migrate && \
    chmod +x /migrate

# 複製編譯後的二進位檔、migration 檔案與 entrypoint 脚本
COPY --from=builder /app/meme-coin .
COPY migrations /migrations
COPY wait-for-it.sh entrypoint.sh /
RUN chmod +x /wait-for-it.sh /entrypoint.sh

EXPOSE 8080
ENTRYPOINT ["/entrypoint.sh"]
