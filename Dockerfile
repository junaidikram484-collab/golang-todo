# =========================
# Stage 1: Build the Go binary
# =========================
FROM golang:1.25-alpine AS builder
# Use a lightweight Alpine-based Go image for building

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker caching
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application as a statically linked binary
# CGO_ENABLED=0 -> disables C bindings, ensures static binary
# GOOS=linux GOARCH=amd64 -> build for Linux 64-bit
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o todo-service ./

# =========================
# Stage 2: Create a minimal runtime image
# =========================
FROM alpine:latest
# Use a tiny Alpine image for runtime to reduce final image size

# Set working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/todo-service .

# Optional: copy any config or docs if needed
# COPY --from=builder /app/docs ./docs

# Expose the port your app runs on
EXPOSE 8080

# Default command to run the binary
CMD ["./todo-service"]