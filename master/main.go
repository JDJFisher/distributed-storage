package main

import (
	"log"
	"net"

	"github.com/JDJFisher/distributed-storage/master/servers"
	"github.com/JDJFisher/distributed-storage/protos"
	"google.golang.org/grpc"
)

func main() {
	// Create a TCP connection on port 6000 for the GRPC server
	listen, err := net.Listen("tcp", ":6000")
	if err != nil {
		log.Fatalf("Failed to open tcp listener... %v", err.Error())
	}

	// Create a GRPC server
	grpcServer := grpc.NewServer()

	// Register Chain service
	chainServer := servers.ChainServer{}
	protos.RegisterChainServer(grpcServer, &chainServer)

	// Register storage service
	storageServer := servers.StorageServer{}
	protos.RegisterStorageServer(grpcServer, &storageServer)

	// Start serving GRPC requests on the open tcp connection
	log.Println("Starting master")
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to start serving the grpc server %v", err.Error())
	}
	defer grpcServer.GracefulStop()
}
