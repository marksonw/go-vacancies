FROM golang:1.22.0

# Set environment variables for CGO
ENV CGO_ENABLED=1
ENV GO111MODULE=on

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN go build -o app .

# Expose the port your application listens on (if applicable)
EXPOSE 8080

# Command to run the application
CMD ["./app"]
