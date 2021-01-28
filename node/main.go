package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

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

	// Create GRPC client
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(grpcHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to the master - %v", err.Error())
	}
	defer conn.Close()

	// Create Chain client
	chainClient := protos.NewChainClient(conn)

	// Repetitively attempt to join the chain
	for {
		_, err := chainClient.Register(context.Background(), &protos.RegisterRequest{Address: os.Getenv("address")})
		if err != nil {
			log.Fatalf("Error joining the chain network - %v", err.Error())
			time.Sleep(5000)
		} else {
			log.Println("Accepted into the chain")
			break
		}
	}

	// TODO: Wait until assigned a role in the chain

	// serve()
}

func serve() {
	// Create a TCP connection for the GRPC server
	listen, err := net.Listen("tcp", os.Getenv("address"))
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
