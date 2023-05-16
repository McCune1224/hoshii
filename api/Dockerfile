FROM golang:1.19-alpine

WORKDIR /app

# Get go mod and sum files
COPY go.mod go.sum ./

# Get dependencies
RUN go mod download

# copy source code into app dir 
COPY . .

# Build app 
RUN go build -o ./bin/main .


# Expose port 8080 to the outside world
EXPOSE 8080

# Run app
CMD ["./bin/main"]
