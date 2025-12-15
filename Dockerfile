FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/myapp


FROM alpine:3.21
WORKDIR /root
COPY --from=builder /app/main .
COPY --from=ghcr.io/kukymbr/goose-docker:3.26.0 /bin/goose /bin/goose
COPY --from=builder /app/migrations ./migrations
CMD ["./main"]