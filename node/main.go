package main

import (
	"context"
	"log"
	"os"

	"github.com/JDJFisher/distributed-storage/protos"
	"google.golang.org/grpc"
)

func main() {
	// Different grpc connection info depending on if it's running in docker or not
	grpcHost := ""
	isDocker := os.Getenv("docker")
	if len(isDocker) == 0 {
		grpcHost = ":6789"
	} else {
		grpcHost = "master:6789"
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
}
