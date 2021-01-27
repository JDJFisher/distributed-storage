package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/JDJFisher/distributed-storage/node/servers"
	"github.com/JDJFisher/distributed-storage/protos"
	"google.golang.org/grpc"
)

func main() {
	// Different grpc connection info depending on if it's running in docker or not
	grpcHost := ":6789"
	if os.Getenv("docker") == "true" {
		grpcHost = "master" + grpcHost
	}

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(grpcHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to the master - %v", err.Error())
	}
	defer conn.Close()

	networkClient := protos.NewNetworkClient(conn)

	response, err := networkClient.JoinNetwork(context.Background(), &protos.NetworkJoinRequest{ServiceName: "server-x"})
	if err != nil {
		log.Fatalf("Error joining the chain network - %v", err.Error())
	}

	log.Printf("Response from master: \n %s", response.Type)

	// TODO: Only serve once accepted in to the chain
	serve()
}

func serve() {
	// Create a TCP connection on port 5000 for the GRPC server
	listen, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Failed to open tcp listener... %v", err.Error())
	}

	// Create a GRPC server
	grpcServer := grpc.NewServer()

	// Register storage service
	storageServer := servers.StorageServer{}
	protos.RegisterStorageServer(grpcServer, &storageServer)

	// Start serving GRPC requests on the open tcp connection
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to start serving the grpc server %v", err.Error())
	}
}
