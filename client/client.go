package main

import (
	"context"
	"log"
	pb "stock-rpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Dial the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials())) // <- for testing only
	if err != nil {
		log.Fatalf("could not connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a new gRPC client
	client := pb.NewStockAnalyserClient(conn)

	// Prepare the request
	req := &pb.Stockrequest{
		Symbol: "IBM",
	}

	// Call GetStockDetail
	resp, err := client.GetStockDetail(context.Background(), req)
	if err != nil {
		log.Fatalf("error calling GetStockDetail: %v", err)
	}

	// Display the response
	log.Printf("Server responded: %s", resp.Message)
}
