# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory in the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . /app

# Download and install any required dependencies
RUN go get -d -v ./...

# Build the Go application
RUN go build -o main .

# Expose a port (if your Go application listens on a specific port)
EXPOSE 8080

# Define the command to run your Go application
CMD ["./main"]