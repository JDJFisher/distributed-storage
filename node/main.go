package main

import (
	"context"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/JDJFisher/distributed-storage/node/servers"
	"github.com/JDJFisher/distributed-storage/protos"
	"google.golang.org/grpc"
)

func main() {
	// Determine port number
	port, err := strconv.Atoi(os.Getenv("port"))
	if err != nil {
		port = 7000
	}

	// Create GRPC client
	var conn *grpc.ClientConn
	conn, err = grpc.Dial("master:6000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to the master - %v", err.Error())
	}
	defer conn.Close()

	// Create Chain client
	chainClient := protos.NewChainClient(conn)

	var neighbours *servers.Neighbours

	// Repetitively attempt to join the chain
	for j := 0; j <= 10; j++ {
		request := &protos.RegisterRequest{Address: os.Getenv("address")}
		response, err := chainClient.Register(context.Background(), request)

		log.Printf("Response to join the network: %v", response.Success)

		//Not allowed to join the network (probably health check related)
		if !response.Success {
			log.Println("Unsuccessful in joining the network, will retry")
			time.Sleep(3 * time.Second)
			log.Printf("Trying to join the network again")
			continue
		}

		//Actual error joining the network
		if err != nil {
			log.Printf("Error joining the chain network - %v (retrying in 3 seconds)", err.Error())
			time.Sleep(3 * time.Second)
			continue
		}

		// Store ...
		neighbours = &servers.Neighbours{
			PredAddress: response.PredAddress,
			SuccAddress: response.SuccAddress,
		}

		log.Println("Accepted into the chain")
		break
	}

	serve(port, conn, neighbours)
}

func serve(port int, conn *grpc.ClientConn, neighbours *servers.Neighbours) {
	// Create a TCP connection for the GRPC server
	listen, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Failed to open tcp listener... %v", err.Error())
	}

	// Create a GRPC server
	grpcServer := grpc.NewServer()

	// Register chain service
	chainServer := servers.NewChainServer(neighbours)
	protos.RegisterChainServer(grpcServer, chainServer)

	// Create a cache
	c := cache.New(cache.NoExpiration, 0)

	// Register storage service
	storageServer := servers.NewStorageServer(neighbours, c)
	protos.RegisterStorageServer(grpcServer, storageServer)

	// Health check client
	healthClient := protos.NewHealthClient(conn)
	go sendHealthCheck(healthClient)

	// Start serving GRPC requests on the open tcp connection
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to start serving the grpc server %v", err.Error())
	}
}

func sendHealthCheck(healthClient protos.HealthClient) {
	for {
		<-time.After(2 * time.Second)
		_, err := healthClient.Alive(context.Background(), &protos.HealthCheckRequest{Address: os.Getenv("address")})
		if err != nil {
			log.Fatalln("Error health checking with the master")
		}
		//log.Printf("Sent health check to master - status: %v", response.Status)
	}

}
