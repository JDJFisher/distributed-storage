package main

import (
	"context"
	"encoding/csv"
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

	// Load dummy requests from seeder
	file, err := os.Open("seeder.csv")
	if err != nil {
		log.Fatal("Unable to read seeder file", err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	lines, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse seeder as CSV", err)
	}

	// Loop over dummy requests
	for i, line := range lines[1:] {
		// Wait ...
		time.Sleep(5000)

		// Determine request type
		requestType := "R"
		if line[1] != "" {
			requestType = "W"
		}

		log.Printf("Dispatching request %v: %v-%v", i, requestType, line[0])

		if requestType == "R" {
			// Fake a read
			_, err = storageClient.Read(context.Background(), &protos.ReadRequest{})
		} else {
			// Fake a write
			_, err = storageClient.Write(context.Background(), &protos.WriteRequest{})
		}

		if err != nil {
			log.Fatalf("Failed request %v: %v", i, err.Error())
		} else {
			log.Printf("Recieved response %v: ...", i)
		}
	}
}
