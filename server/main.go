package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ritikkoul0/stock-rpc/logger"
	"github.com/ritikkoul0/stock-rpc/server/operations/overview"

	"github.com/ritikkoul0/stock-rpc/server/utils"

	pb "github.com/ritikkoul0/stock-rpc/proto"

	"github.com/ritikkoul0/stock-rpc/database"

	"google.golang.org/grpc"
)

// stockAnalyserServer implements the stockrpc.stockAnalyser gRPC service.
type stockAnalyserServer struct {
	pb.UnimplementedStockAnalyserServer
}

// GetStockDetail is the implementation of the RPC method.
func (s *stockAnalyserServer) GetStockDetail(ctx context.Context, req *pb.Stockrequest) (*pb.StockResponse, error) {
	log.Printf("Received request for symbol: %s", req.Symbol)
	client := &http.Client{}
	response, err := client.Get("https://www.alphavantage.co/query?function=OVERVIEW&symbol=" + req.Symbol + "&apikey=J0FU1KW89DFXSHQH")
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	var company overview.Overview
	if err := json.Unmarshal(body, &company); err != nil {
		fmt.Println("Error unmarshalling json:", err)
		return nil, err
	}
	error := database.Insertoverview(ctx, company)
	if error != nil {
		resp := &pb.StockResponse{
			Message: "Failed added overview for stock " + req.Symbol,
		}
		return resp, nil
	}

	resp := &pb.StockResponse{
		Message: "Successful added overview for stock " + req.Symbol,
	}
	return resp, nil
}

var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGINT}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	utils.UpdateVariables()
	logger.NewLogger("development")
	appContext, stopSignals := signal.NotifyContext(context.Background(), shutdownSignals...)
	defer stopSignals()
	logger.Info("Initializing database connection...")
	err = database.InitializeConnection(appContext, utils.Config)
	if err != nil {
		logger.Fatalf("Failed to initialize database connection: %v", err)
	}
	logger.Infof("Database connection initialized successfully")

	grpcServer := grpc.NewServer()
	pb.RegisterStockAnalyserServer(grpcServer, &stockAnalyserServer{})

	fmt.Println("gRPC server listening on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
