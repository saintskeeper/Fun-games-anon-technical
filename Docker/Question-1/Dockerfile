# Build stage
FROM golang:latest AS builder

# Set the working directory
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application
RUN go build -o app

# Final stage
FROM alpine:latest

# Copy the binary from the build stage into the final image
COPY --from=builder /app/app /usr/local/bin/app

# Set the entry point for the container
ENTRYPOINT ["/usr/local/bin/app"]
