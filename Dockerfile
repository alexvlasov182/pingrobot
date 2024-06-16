# Step 1: Build the Go application
FROM golang:1.22.1 AS builder
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o pingrobot main.go

# Step 2: Create a lightweight container to run the Go application
FROM alpine:latest
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/pingrobot .

# Expose the application's port
EXPOSE 8080

# Run the binary
CMD ["./pingrobot"]
