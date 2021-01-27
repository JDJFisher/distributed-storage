package main

import (
	"context"
	"log"
	"os"

	"github.com/JDJFisher/distributed-storage/protos"
	"google.golang.org/grpc"
)

func main() {
	// Different grpc connection info depending on if it's running in docker or not
	grpcHost := ""
	isDocker := os.Getenv("docker")
	if len(isDocker) == 0 {
		grpcHost = ":6789"
	} else {
		grpcHost = "master:6789"
	}

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(grpcHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to the master - %v", err.Error())
	}
	defer conn.Close()

	storageClient := protos.NewStorageClient(conn)

	response, err := storageClient.Read(context.Background(), &protos.ReadRequest{Message: "foo"})
	if err != nil {
		log.Fatalf("Error reading from master - %v", err.Error())
	}

	log.Printf("Read response from master: \n %s", response)
}
