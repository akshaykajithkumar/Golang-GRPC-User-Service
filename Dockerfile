# Start from the official golang image with the desired version
FROM golang:1.21.4

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o server cmd/server/main.go

# Expose port 50051 to the outside world
EXPOSE 50051

# Command to run the executable
CMD ["./server"]

