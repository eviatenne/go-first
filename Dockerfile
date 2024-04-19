# Use a Golang base image
FROM golang:1.22.2

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN go build -o myapp .

# Expose the port your Go application listens on
EXPOSE 8080

# Command to run the executable
CMD ["./myapp"]
