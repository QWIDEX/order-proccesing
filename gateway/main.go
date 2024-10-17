package main

import (
	"fmt"
	pb "github.com/QWIDEX/order-proccesing/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	ordersConn, err := grpc.NewClient("localhost:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("cannot connect to orders service")
		return
	}
	defer ordersConn.Close()

	log.Println("Connected to orders service successfully")

	orderClient := pb.NewOrderServiceClient(ordersConn)
	server := NewServer(orderClient)

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
