# Use the official Golang 1.20 image as the base image
FROM golang:1.21-alpine3.18

# Set the working directory inside the container
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN apk add --no-cache make

# Copy the rest of the application code to the container
# first dot represents the root of the folder
# second dot represents the current working directory inside the image where the files are copied to here /app
COPY . .

# Expose port 8080 for the backend server
EXPOSE 8080

# Build the application
# RUN go build -o /main

CMD ["./startup.sh"]

# Start the backend server
# CMD ["/main"]
