# Postgres Database server setup
FROM postgres:14-alpine AS backend-database

ENV POSTGRES_PASSWORD postingdata

ENV POSTGRES_DB cshare

ENV POSTGRES_PORT 5432

EXPOSE 5432
# init.sql is used here because this container is also used for testing pusposes.
# idely, this file should be removed when the container is used for production. It should be in the docker-compose.yml file.
COPY models/init.sql /docker-entrypoint-initdb.d/

# Start from golang base image
FROM golang:alpine as builder

ENV GO111MODULE=on

# Add Maintainer info
# LABEL maintainer="Utkarsh Chourasia<utkarshchourasia.in>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

# Start a new stage from scratch
FROM alpine:latest AS backend-server
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/server .
COPY --from=builder /app/.env .       

# Expose port 5675 to the outside world
EXPOSE 5675

#Command to run the executable
CMD ["./server"]
