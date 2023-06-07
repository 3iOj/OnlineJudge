# Use the official Golang 1.16 image as the base image
FROM golang:1.18

# Set the working directory inside the container
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the rest of the application code to the container
COPY . .

# Expose port 8080 for the backend server
EXPOSE 8080

# Build the application
# RUN go build -o /main

CMD ["air", "-c", ".air.toml"]

# Start the backend server
# CMD ["/main"]