# Build Stage
FROM golang:1.23 AS builder

WORKDIR /app

# Copy Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go application
RUN go build -o /app/main .

# Final Runtime Stage
FROM alpine:latest

WORKDIR /root/

# Install libc (needed for Go binaries on Alpine)
RUN apk add --no-cache libc6-compat

# Copy the built binary from the builder stage to the root folder
COPY --from=builder /app/main /main

# Ensure execution permissions
RUN chmod +x /main

# Expose the application port
EXPOSE 8080

COPY .env .env

# Run the application
CMD ["/main"]

