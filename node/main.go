package main

import (
	"context"
	"log"

	"github.com/JDJFisher/distributed-storage/protos"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("master:6789", grpc.WithInsecure())
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
