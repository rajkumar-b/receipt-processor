# Docker file based on helloworld app from 
# https://github.com/GoogleCloudPlatform/golang-samples/blob/main/run/helloworld/Dockerfile

# Use the offical golang image to create a binary.
FROM golang:1.21-bookworm as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . .

# Build the binary.
RUN go build -v -o receipt-processor

# Use the official Debian slim image for a lean production container.
FROM debian:bookworm-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/receipt-processor /app/receipt-processor

# Expose port 8080
EXPOSE 8080

# Run the web service on container startup.
CMD ["/app/receipt-processor"]
