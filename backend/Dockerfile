# syntax=docker/dockerfile:1
FROM golang:1.17.1-alpine

# Create a directory for the application.
WORKDIR /maxpondus

# Download dependencies.
COPY go.mod .
RUN go mod download

# Copy project files and build the application.
COPY . .
RUN go build -o /bin/mxp ./cmd/server

# Specify the executable to run in the container.
CMD ["/bin/mxp"]

# Tell Docker about the port used by the application.
EXPOSE 8080