
# Build the images
SERVICE=""
build:
	docker-compose build $(SERVICE)


# Run the chain services
serve:
	docker-compose up


# Execute a request on the chain
request:
	docker-compose run --rm -e OP=$(OP) -e KEY=$(KEY) -e VALUE=$(VALUE) client


# Regenerate all the proto files
PROTO=*
proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/$(PROTO).proto


# Delete generated grpc sources
clean:
	rm -f protos/*.pb.go


#Run tests for all of the code
test:
	go test -v ./...