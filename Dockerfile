# Step 1: Build the Go binary
FROM golang:1.24-alpine AS builder

# Install git for go modules
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the project
COPY . .

# Build the binary
RUN go build -o subs ./main.go

# Step 2: Create a small image to run the binary
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/subs .

# Copy .env file
COPY .env .

# Expose the port
EXPOSE 7777

# Run the binary
CMD ["./subs"]
