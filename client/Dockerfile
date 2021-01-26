FROM golang:1.14-alpine

RUN mkdir /app

# Copy everything in here into the app directory
ADD . /app

WORKDIR /app

# Build the go project into a binary
RUN go build -o main .

# Run the freshly build main executable
CMD ["./main"]