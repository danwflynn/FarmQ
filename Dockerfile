# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG SERVICE
RUN go build -o app ./cmd/${SERVICE}

# Runtime stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/app .

CMD ["./app"]
