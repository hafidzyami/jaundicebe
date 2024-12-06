# Stage 1: Build the Go binary
FROM golang:1.20-alpine as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go modules manifests to the container
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code to the container
COPY . .

# Build the Go application
RUN go build -o main .

# Stage 2: Create the final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the Go binary from the builder stage
COPY --from=builder /app/main .

# Copy the .env file to the container (if needed)
COPY .env .env

# Expose the port the app will run on
EXPOSE 3000

# Command to run the app
CMD ["./main"]
