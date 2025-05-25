# Build stage
FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

# Copy the source code
COPY . .

# Download dependencies
RUN go mod tidy

# Build the application
RUN GOOS=linux GOARCH=amd64 go build -o /tmp/hello .

# Final stage
FROM alpine:latest

# Set working directory
WORKDIR /newapp

# Copy the binary from builder
COPY --from=builder /tmp/hello .

# Expose port
EXPOSE 3000

# Command to run the executable
ENTRYPOINT ["/newapp/hello"]

