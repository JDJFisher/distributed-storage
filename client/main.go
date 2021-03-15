package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/JDJFisher/distributed-storage/protos"
	"google.golang.org/grpc"
)

func main() {
	// Different grpc connection info depending on if it's running in docker or not
	grpcHost := ":6000"
	if os.Getenv("docker") == "true" {
		grpcHost = "master" + grpcHost
	}

	// Open a connection to the master service
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(grpcHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to the master - %v", err.Error())
	}
	defer conn.Close()

	// Create storage client
	storageClient := protos.NewStorageClient(conn)

	// Fake requests
	fake(storageClient)
}

func fake(storageClient protos.StorageClient) {

	// Declare dummy requests
	var dummyRequests = [...]interface{}{
		&protos.WriteRequest{Key: "alpha", Value: "foo"},
		&protos.ReadRequest{Key: "alpha"},
		&protos.WriteRequest{Key: "alpha", Value: "bar"},
		&protos.WriteRequest{Key: "beta", Value: "baz"},
		&protos.ReadRequest{Key: "alpha"},
		&protos.ReadRequest{Key: "beta"},
	}

	time.Sleep(2 * time.Second)

	// Loop over dummy requests
	for i, request := range dummyRequests {
		// Wait ...
		// time.Sleep(2 * time.Second)

		switch request := request.(type) {
		// Panic
		default:
			log.Panicln("Unexpected request type")

		// Fake a read
		case *protos.ReadRequest:
			log.Printf("Dispatching read %v: %v", i, request.Key)

			response, err := storageClient.Read(context.Background(), request)

			if err != nil {
				log.Fatalf("Failed read %v: %v", i, err.Error())
			} else {
				log.Printf("Recieved response %v: %v", i, response.Value)
			}

		// Fake a write
		case *protos.WriteRequest:
			log.Printf("Dispatching write %v: %v->%v", i, request.Key, request.Value)

			_, err := storageClient.Write(context.Background(), request)

			if err != nil {
				log.Fatalf("Failed write %v: %v", i, err.Error())
			} else {
				log.Printf("Recieved response %v", i)
			}
		}
	}
}
