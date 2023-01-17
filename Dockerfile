# Start from base image
FROM golang:1.18-alpine as builder
LABEL MAINTENAIR allan.nava@hiway.media
# Set the current working directory inside the container
WORKDIR /app
# Copy go mod and sum files
COPY go.mod go.sum ./
# Download all dependencies
RUN go mod download
# Copy source from current directory to working directory
COPY . .
# Build the application
# Produce binary named main
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -o  main cmd/todo/todo-api.go
#################
#FROM phusion/baseimage:focal-1.2.0
FROM phusion/baseimage:jammy-1.0.1
#
COPY --from=build /app/main .
#
EXPOSE 8080
CMD ["./main"]
#
