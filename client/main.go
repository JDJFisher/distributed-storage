package main

import (
	"context"
	"log"

	"github.com/JDJFisher/distributed-storage/protos"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("master:6789", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to the master - %v", err.Error())
	}
	defer conn.Close()

	storageClient := protos.NewStorageClient(conn)

	response, err := storageClient.read(context.Background(), &protos.ReadRequest{message: "foo"})
	if err != nil {
		log.Fatalf("Error reading from master - %v", err.Error())
	}

	log.Printf("Read response from master: \n %s", response.Type)
}
