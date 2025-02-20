# 使用 golang 基礎映像
FROM golang:1.23 AS builder

# 設定工作目錄
WORKDIR /app

# 設置環境變數
ENV SERVER_ADDR=localhost:8889

# 複製並編譯
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o myapp ./main

# 使用 alpine 基礎映像
FROM alpine:latest

# 設定工作目錄
WORKDIR /app

# 複製編譯好的程式
COPY --from=builder /app/myapp .

# 設定執行命令
ENTRYPOINT ["./myapp"]
