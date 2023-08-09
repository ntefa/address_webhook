# Use the official Go image as the base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the rest of the application code
COPY . .
#RUN go mod tidy

# Build the Go application
RUN go build -o webhook-receiver

# Expose the port your application listens on
EXPOSE 8080

# Command to run your application
CMD ["./webhook-receiver"]
