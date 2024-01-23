# Use an official lightweight Go image
FROM golang:alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 7777
EXPOSE 7777

# Set the entry point to the application executable
CMD ["./main"]