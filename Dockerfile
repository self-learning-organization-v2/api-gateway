# Stage 1: Build the Go binary
FROM golang:1.24.5-alpine AS builder

RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o api-gateway main.go

# Stage 2: Minimal image
FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/api-gateway .
COPY config.json .
EXPOSE 3000
CMD ["./api-gateway"]
