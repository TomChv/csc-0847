###
# Stage: builder
# Install dependencies and build backend binary.
# Output:
#  /app/binary
###
FROM golang:1.20-alpine AS builder

# Define working directory.
WORKDIR /app

# Install build dependencies.
RUN apk add build-base

# Install Golang dependencies.
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

# Copy source code.
COPY . .

# Build nightclubpoll binary.
RUN go build -o /app/backend ./main.go

###
# Stage: app
# Expose nightclubpoll binary entrypoint in a light image.
###
FROM alpine AS app

# Define workding directory
WORKDIR /app

# Copy binary from builder stage to alpine image.
COPY --from=builder /app/backend /app/backend

# Set entrypoint to nightclubpoll binary.
ENTRYPOINT ["/app/backend"]