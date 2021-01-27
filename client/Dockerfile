# Using an alpine image because it removes all the bs in linux which makes the images so large
# literally just a fresh install with only things needed for go to work
FROM golang:1.15-alpine

RUN mkdir /app

# Copy everything in here into the app directory
ADD . /app

WORKDIR /app/client

# Download the depdencies of the project
RUN go mod download

# # Build the go project into a binary
RUN go build main.go

# Expose the GRPC port we'll be using
EXPOSE 6789

# #Run the freshly build main executable
CMD ["./main"]