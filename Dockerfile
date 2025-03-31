FROM golang:latest AS builder

# Set the working directory inside the container
WORKDIR /api

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o main .

# Use scratch to reduce image size
FROM scratch

# Copy certificates and 
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Copy executable
COPY --from=builder /api/main /main

ENV APP_PORT=8080
# Expose the port that the application runs on
EXPOSE 8080

# Command to run the application
ENTRYPOINT ["/main"]
