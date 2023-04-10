# Stage 1: Build the Golang binary
FROM golang:1.20-alpine AS builder

WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Golang binary with static linking
RUN go build cmd/cancel-transaction-worker/main.go

# Stage 2: Build the final Docker image
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the build stage into the final image
COPY --from=builder /app/main ./

USER root

RUN chmod +x main

# Set the command to run when the container starts
CMD ["./main"]
