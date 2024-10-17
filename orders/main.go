package main

import (
	"context"
	"github.com/QWIDEX/order-proccesing/common"
	"google.golang.org/grpc"
	"log"
	"net"

	_ "github.com/joho/godotenv/autoload"
)

var (
	grpcAddr = common.GetEnv("grpcAddr", ":3000")
)

func main() {
	server := grpc.NewServer()

	db := NewOrdersDB(context.Background())
	s := NewService(db)

	NewGRPCHandler(server, s)

	conn, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatal("failed to listen ", err)
	}
	defer conn.Close()

	if err := server.Serve(conn); err != nil {
		log.Fatal(err.Error())
	}
}
