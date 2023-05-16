# Start from a minimal Alpine Linux image with the latest version of Go
FROM golang:1.19-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application
RUN go build -o /app/server ./cmd

# Expose the port that the server will listen on
EXPOSE 8080

# Start the server when the container starts
CMD ["/app/server"]
