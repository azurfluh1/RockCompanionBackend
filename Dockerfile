# Build stage
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /usr/local/bin/server

# Run stage
FROM debian:bookworm-slim

COPY --from=builder /usr/local/bin/server /usr/local/bin/server

RUN chmod +x /usr/local/bin/server

EXPOSE 8080

CMD ["/usr/local/bin/server"]