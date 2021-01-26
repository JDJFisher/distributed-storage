run:
	docker-compose up

# Compile a specific proto file - usage: "make argument=health proto"
proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative **/$(argument).proto

# Regenerate all the proto files
proto-all:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative **/*.proto

test:
	pytest

clean:
	docker-compose rm -f