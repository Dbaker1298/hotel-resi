FROM golang:1.21.4-bullseye

# Set the working directory
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY go.mod go.sum ./

# Download all the dependencies
RUN go mod tidy

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./main"]
