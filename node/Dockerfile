# Using an alpine image because it removes all the bs in linux which makes the images so large
# literally just a fresh install with only things needed for go to work
FROM golang:1.15-alpine as build
RUN apk add git

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY protos .

ENV GO111MODULE=on
RUN go mod download

COPY . .

RUN go build ./node/main.go

# <- Second step to build minimal image
FROM golang:1.15-alpine
COPY --from=build /app/main main
RUN ["chmod", "+x", "./main"]
CMD ["./main"]
