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
RUN go build .

# Expose port 8080
EXPOSE 8080

# Run the web service on container startup.
CMD ["./receipt-processor"]
