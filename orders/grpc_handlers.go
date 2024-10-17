package main

import (
	"context"
	pb "github.com/QWIDEX/order-proccesing/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer

	service Service
}

func NewGRPCHandler(grpcServer *grpc.Server, s Service) {
	handler := grpcHandler{service: s}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h grpcHandler) CreateOrder(context.Context, *pb.CreateOrderReq) (*pb.Order, error) {
	return nil, status.Errorf(codes.OK, "method CreateOrder implemented")
}
