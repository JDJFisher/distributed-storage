FROM golang:1.15-alpine

RUN mkdir /app

# Copy everything in here into the app directory
ADD . /app

WORKDIR /app

# Download the depdencies of the project
RUN go mod download

# Build the go project into a binary
RUN go build -o main .

# Run the freshly build main executable
CMD ["./main"]