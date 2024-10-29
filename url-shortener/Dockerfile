# Stage 1: Build the Go application
FROM golang:1.18-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files for dependency resolution
COPY go.mod go.sum ./

# Download dependencies early to leverage Docker caching
RUN go mod download

# Copy the rest of the application code
COPY . .

# Specify the path to the main application file in the cmd directory
RUN go build -o urlshortener ./cmd/main.go

# Stage 2: Run the application in a minimal container
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/urlshortener .

# Copy config folder
COPY ./conf /root/conf

# Expose the port the application will listen on (change if needed)
EXPOSE 80

# Command to run the application
CMD ["./urlshortener"]
