# Stage 1: Build the Go application
FROM golang:1.22 AS build

WORKDIR /app

# Copy Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o main ./cmd

# Copy the .env file
COPY .env .env
# Set environment variables
ENV $(cat .env | xargs)

# Set the entrypoint to run the Go application
CMD ["./main"]

