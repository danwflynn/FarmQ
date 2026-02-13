# Start from the official Go image
FROM golang:1.25-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first for caching dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the binary
RUN go build -o farmq ./cmd/farmq

# Expose the port your API uses
EXPOSE 8080

# Run the binary
CMD ["./farmq"]
