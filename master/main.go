package main

import (
	"log"
	"net"

	"github.com/JDJFisher/distributed-storage/master/servers"
	"github.com/JDJFisher/distributed-storage/protos"
	"google.golang.org/grpc"
)

func main() {
	// Create a TCP connection on port 6789 for the GRPC server
	listen, err := net.Listen("tcp", ":6789")
	if err != nil {
		log.Fatalf("Failed to open tcp listener... %v", err.Error())
	}

	// Create a GRPC server
	grpcServer := grpc.NewServer()

	// Register Network service
	networkServer := servers.NetworkServer{}
	protos.RegisterNetworkServer(grpcServer, &networkServer)

	// Register storage service
	storageServer := servers.StorageServer{}
	protos.RegisterStorageServer(grpcServer, &storageServer)

	// Start serving GRPC requests on the open tcp connection
	log.Println("[MASTER] Starting master.... GRPC serving on port: 6789")
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to start serving the grpc server %v", err.Error())
	}
	defer grpcServer.GracefulStop()
}
