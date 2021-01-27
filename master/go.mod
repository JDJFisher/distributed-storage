module github.com/JDJFisher/distributed-storage/master

go 1.15

require (
	github.com/JDJFisher/distributed-storage/protos v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.35.0
)

replace github.com/JDJFisher/distributed-storage/protos => ../protos
