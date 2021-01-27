package main

import (
	"context"
	"encoding/csv"
	"log"
	"os"

	"github.com/JDJFisher/distributed-storage/protos"
	"google.golang.org/grpc"
)

func main() {
	// Different grpc connection info depending on if it's running in docker or not
	grpcHost := ":6789"
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

		// TODO: Sleep for a few seconds

		log.Println("Dispatching request", i, ":", line)

		// TODO: Execute a Write if line has a value value otherwise do a read
	}

	response, err := storageClient.Read(context.Background(), &protos.ReadRequest{})
	if err != nil {
		log.Fatalf("Error reading from master - %v", err.Error())
	} else {
		log.Printf("Response from master: \n %s", response)
	}
}
