package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	pb "wjuniorbh92/go-grpc-poc/protos"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type SaveShareServer struct {
	pb.UnimplementedSaveSharePriceServer
}

func (s *SaveShareServer) StoreStockDatabase(stream pb.SaveSharePrice_StoreStockDatabaseServer) error {
	startTime := time.Now()
	var stockData []*pb.StockPrice
	for {
		msg, err := stream.Recv()
		stockData = append(stockData, msg)
		if err == io.EOF {
			endTime := time.Now()
			fmt.Printf("Time to Received the data: %f seconds \n", float64(endTime.Sub(startTime).Seconds()))
			return stream.SendAndClose(&pb.PriceResponse{
				Symbol: stockData[0].Symbol,
			})
		}
		if err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSaveSharePriceServer(s, &SaveShareServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
