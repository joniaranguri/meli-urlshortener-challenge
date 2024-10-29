# Start from the official Golang image
FROM golang:1.18-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -o id_generator .

# Start a new, minimal image for the final container
FROM alpine:latest

# Install necessary libraries
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/id_generator .

# Copy config folder
COPY ./conf /root/conf

# Command to run the executable
CMD ["./id_generator"]
