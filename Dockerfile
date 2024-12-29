FROM golang:1.23 AS builder
LABEL authors="Vincent Han <switchover@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./

# Download all dependencies (caching layer)
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o egovchecker .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ./build_with_flags.sh -o egovchecker

FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/egovchecker .

# Run the binary
ENTRYPOINT ["./egovchecker"]
