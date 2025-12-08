FROM golang:1.24-alpine AS BUILDER

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/myapp

FROM alpine:3.21
WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
