FROM golang:1.23 AS builder
LABEL authors="Vincent Han <switchover@gmail.com>"

# Timezone
ENV TZ=Asia/Seoul

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

# Timezone
ENV TZ=Asia/Seoul

# Set the working directory.
# - The directory to be inspected must be mounted on the host via volume.
WORKDIR /target

# Copy the binary from the builder stage
COPY --from=builder /app/egovchecker /app/egovchecker

# Run the binary
ENTRYPOINT ["/app/egovchecker"]
