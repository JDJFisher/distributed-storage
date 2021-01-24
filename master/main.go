package main

import (
	"log"
	"net"

	"github.com/JDJFisher/primary-backup/master/network"
	"google.golang.org/grpc"
)

func main() {

	//Create a TCP  connection on port 6789 for the GRPC server
	listen, err := net.Listen("tcp", ":6789")
	if err != nil {
		log.Fatalf("Failed to open tcp listener... %v", err.Error())
	}

	networkServer := network.Server{}

	//Create the new GRPC server
	grpcServer := grpc.NewServer()

	network.RegisterNetworkServer(grpcServer, &networkServer)

	//Start serving GRPC requests on the open tcp connection
	log.Println("[MASTER] Starting master.... GRPC serving on port: 6789")
	err = grpcServer.Serve(listen)

	if err != nil {
		log.Fatalf("Failed to start serving the grpc server %v", err.Error())
	}

}
