FROM golang:1.21.3-alpine3.18

# Set working directory
WORKDIR /app

# Copy files into container
COPY . ./

# Download dependencies
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-go-server

# Expose HTTP port
EXPOSE 9990

# Run
CMD ["/docker-go-server"]