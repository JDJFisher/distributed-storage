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
	// Collect args from the environment
	key := os.Getenv("KEY")
	value := os.Getenv("VALUE")

	// Exit if no key was specified
	if key == "" {
		return
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

	// Send a request
	if value != "" {
		sendWriteRequest(client, &protos.WriteRequest{Key: key, Value: value})
	} else {
		sendReadRequest(client, &protos.ReadRequest{Key: key})
	}
}

func sendReadRequest(client protos.StorageClient, request *protos.ReadRequest) {
	log.Println("Dispatching read:", request.Key)

	response, err := client.Read(context.Background(), request)

	if err != nil {
		log.Fatalln("Failed read:", err.Error())
	} else if response.Value == "" {
		log.Println("Recieved empty value")
	} else if response.Value == "" {
		log.Println("Recieved value:", response.Value)
	}
}

func sendWriteRequest(client protos.StorageClient, request *protos.WriteRequest) {
	log.Printf("Dispatching write: %v->%v", request.Key, request.Value)

	_, err := client.Write(context.Background(), request)

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
