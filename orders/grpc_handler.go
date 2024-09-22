package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/ad-bak/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}

func NewGRPCHandler(grpcServer *grpc.Server) {
	handler := &grpcHandler{}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {
	return nil, fmt.Errorf("test error !")
	log.Println("New order received!")
	log.Printf("Order created: %v", p)
	o := &pb.Order{
		ID: "42",
	}
	return o, nil
}
