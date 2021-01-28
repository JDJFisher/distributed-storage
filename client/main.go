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
		time.Sleep(5000)
		log.Println("Dispatching request", i, ":", line)

		if line[1] == "" {
			// Fake a read
			response, err := storageClient.Read(context.Background(), &protos.ReadRequest{})
			if err != nil {
				log.Fatalln("Error reading", err.Error())
			} else {
				log.Println("Read Response", response)
			}
		} else {
			// Fake a write
			response, err := storageClient.Write(context.Background(), &protos.WriteRequest{})
			if err != nil {
				log.Fatalln("Error writing", err.Error())
			} else {
				log.Println("Read Response", response)
			}
		}
	}
}
