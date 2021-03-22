package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/JDJFisher/distributed-storage/protos"
	"google.golang.org/grpc"
)

func main() {
	log.SetOutput(os.Stdout)

	// Parse command args
	args := os.Args[1:]
	if len(args) < 2 {
		log.Println("Too few args")
		return
	}
	op := args[0]
	key := args[1]
	value := ""
	if len(args) >= 3 {
		value = args[2]
	}

	// Open a connection to the master service
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(os.Getenv("host"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to the master - %v", err.Error())
	}
	defer conn.Close()

	// Create storage client
	client := protos.NewStorageClient(conn)

	// Make a request
	switch strings.ToUpper(op) {
	case "READ":
		sendReadRequest(client, &protos.ReadRequest{Key: key})
	case "WRITE":
		sendWriteRequest(client, &protos.WriteRequest{Key: key, Value: value})
	default:
		log.Println("Invalid operation: ", op)
	}
}

func sendReadRequest(client protos.StorageClient, request *protos.ReadRequest) {
	log.Println("Requesting read:", request.Key)

	// Execute
	response, err := client.Read(context.Background(), request)

	// Display response
	if err != nil {
		log.Fatalln("Failed read:", err.Error())
	} else if response.Value == "" {
		log.Println("Recieved empty value")
	} else {
		log.Println("Recieved value:", response.Value)
	}
}

func sendWriteRequest(client protos.StorageClient, request *protos.WriteRequest) {
	if request.Value == "" {
		log.Printf("Requesting empty write: %v", request.Key)
	} else {
		log.Printf("Requesting write: %v->%v", request.Key, request.Value)
	}

	// Execute
	_, err := client.Write(context.Background(), request)

	// Display response
	if err != nil {
		log.Fatalln("Failed write:", err.Error())
	} else {
		log.Println("Write persisted")
	}
}

// -----------------------------------------------------------------------------------

// Declare dummy requests
var dummyRequests = []interface{}{
	&protos.WriteRequest{Key: "alpha", Value: "foo"},
	&protos.ReadRequest{Key: "alpha"},
	&protos.WriteRequest{Key: "alpha", Value: "bar"},
	&protos.WriteRequest{Key: "beta", Value: "baz"},
	&protos.ReadRequest{Key: "alpha"},
	&protos.ReadRequest{Key: "beta"},
}

func fakeRequests(client protos.StorageClient, requests []interface{}) {
	// Loop over dummy requests
	for i, request := range requests {
		// Wait ...
		time.Sleep(time.Second)
		log.Println("Faking request:", i)

		switch request := request.(type) {

		// Fake a read
		case *protos.ReadRequest:
			sendReadRequest(client, request)

		// Fake a write
		case *protos.WriteRequest:
			sendWriteRequest(client, request)
		}
	}
}
