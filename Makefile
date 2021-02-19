run:
	docker-compose up


# Regenerate all the proto files
PROTO=*
proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/$(PROTO).proto

clean:
	rm -f protos/*.pb.go
	# docker-compose rm -f