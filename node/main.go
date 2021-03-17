package main

import (
	"context"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/JDJFisher/distributed-storage/node/servers"
	"github.com/JDJFisher/distributed-storage/protos"
	"github.com/patrickmn/go-cache"
	"google.golang.org/grpc"
)

func main() {
	// Determine port number
	port, err := strconv.Atoi(os.Getenv("port"))
	if err != nil {
		port = 7000
	}

	// Create GRPC client
	masterConn, err := grpc.Dial(os.Getenv("host"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to the master - %v", err.Error())
	}
	defer masterConn.Close()

	// Create Chain client
	chainClient := protos.NewChainClient(masterConn)

	// Create a cache
	c := cache.New(cache.NoExpiration, 0)

	// Repetitively attempt to join the chain
	for j := 0; j <= 10; j++ {
		response, err := chainClient.GetTail(context.Background(), &protos.TailRequest{})

		// Not allowed to join the network
		if err != nil {
			log.Printf("Error joining the chain network - %v (retrying in 3 seconds)", err.Error())
			time.Sleep(3 * time.Second)
			continue
		}

		tailAddress := response.Address
		log.Printf("Fetched current chain tail - %v", tailAddress)

		if tailAddress != "" {
			//
			tailConn, err := grpc.Dial(response.Address, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("Error connecting to the tail to grab the data - %v", err.Error())
			}
			defer tailConn.Close()

			// Create Storage client
			tailClient := protos.NewStorageClient(tailConn)

			// Stream the data
			stream, err := tailClient.GetTailData(context.Background(), &protos.RequestData{})
			if err != nil {
				log.Fatalf("Error getting the tail data - %v", err.Error())
			}
			for {
				log.Println("Running client loop")
				item, err := stream.Recv()
				if err == io.EOF {
					log.Println("Received all data from the tail")
					break
				} else if err != nil {
					log.Fatalf("Error receiving key data from tail whilst getting data - %v", err.Error())
				}
				// Add the key value pair to the cache locally.
				log.Printf("Writing %s: %s", item.Key, item.Value)
				c.Set(item.Key, item.Value, cache.NoExpiration)
			}
		}

		// Join the chain
		request := &protos.JoinRequest{
			Address:     os.Getenv("address"),
			TailAddress: tailAddress,
		}
		_, err = chainClient.Join(context.Background(), request)

		if err != nil {
			log.Println("find in folder this error because its secretive")
			continue
		}

		log.Println("Accepted into the chain")

		neighbours := &servers.Neighbours{
			PredAddress: tailAddress,
			SuccAddress: "",
		}

		serve(port, masterConn, neighbours, c)
	}
}

func serve(port int, masterConn *grpc.ClientConn, neighbours *servers.Neighbours, c *cache.Cache) {
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

	// Register storage service
	storageServer := servers.NewStorageServer(neighbours, c)
	protos.RegisterStorageServer(grpcServer, storageServer)

	// Health check client
	healthClient := protos.NewHealthClient(masterConn)
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
