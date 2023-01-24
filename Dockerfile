# Start from golang base image
FROM golang:alpine as builder

# Add Maintainer Info
LABEL maintainer="Utkarsh Chourasia <5.utkarshchourasia@gmail.com"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download 

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest AS production
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Port to the outside world
EXPOSE 5675

#Command to run the executable
CMD ["./main"]

# Postgres Database setup
FROM postgres:14-alpine AS database

# Copy the database schema
COPY models/init.sql /docker-entrypoint-initdb.d/

# Expose the port for connection.
EXPOSE 5432