FROM golang:1.22.5

# Set correct working direcotry inside the container
WORKDIR /app

# Install Air for live reloading
RUN go install github.com/air-verse/air@latest

# Copy go.mod go.sum
COPY go.mod go.sum ./

# Download all dependenices
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# # Install necessary build dependencies
# RUN apt-get update && \
#     apt-get install -y gcc && \
#     apt-get install -y musl-dev && \
#     rm -rf /var/lib/apt/lists/*

# Enable cgo for Go builds
# ENV CGO_ENABLED=1

# Expose the  port the app runs on
EXPOSE 8000

# Run the application
CMD ["air", "-c", ".air.toml"]
