# Pull the official Golang image with golang version 1.22
FROM golang:1.22

# Set/create working directory inside the container
WORKDIR /app

# Copy go.mod go.sum files into working directory in container
COPY go.mod go.sum ./

# Download Dependencies from go mod
RUN go mod download

# Copy the rest of local folder into the working directory in container
COPY . . 

# Build the Go application and the output are binary within the 'bin' working directory (/app/bin)
RUN go build -o bin main.go

# Set the binary as the command to run when the container starts
ENTRYPOINT [ "/app/bin" ]