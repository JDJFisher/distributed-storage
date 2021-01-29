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
		time.Sleep(2 * time.Second)

		// Determine request type
		requestType := "R"
		if line[1] != "" {
			requestType = "W"
		}

		if requestType == "R" {
			log.Printf("Dispatching read %v: %v", i, line[0])

			// Fake a read
			request := protos.ReadRequest{Key: line[0]}
			response, err := storageClient.Read(context.Background(), &request)

			if err != nil {
				log.Fatalf("Failed read %v: %v", i, err.Error())
			} else {
				log.Printf("Recieved response %v: %v", i, response.Value)
			}
		} else {
			log.Printf("Dispatching write %v: %v->%v", i, line[0], line[1])

			// Fake a write
			request := protos.WriteRequest{Key: line[0], Value: line[1]}
			response, err := storageClient.Write(context.Background(), &request)

			if err != nil {
				log.Fatalf("Failed write %v: %v", i, err.Error())
			} else {
				log.Printf("Recieved response %v: %v", i, response.Value)
			}
		}

	}
}
