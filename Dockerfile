FROM golang:1.19-alpine AS build

# Set up the work directory
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod download

# Copy the application source code
COPY pkg ./pkg
COPY cmd ./cmd

# Build the application
RUN go build -o /app/mock ./cmd/main.go

# Use a minimal base image for the final container (is alpine still the best choice?)
FROM alpine:latest

# Copy the built executable from the build stage
COPY --from=build /app/mock /app/mock

# Set the working directory
WORKDIR /app

# Run the application
CMD ["./mock"]