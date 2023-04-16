# Use the official Go image as a base image
FROM golang:1.20.3

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application
RUN go build -o sample-api cmd/main.go

# Expose the port on which the application will run
EXPOSE 8080

# Run the application
CMD ["./sample-api"]
