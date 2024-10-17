package main

import (
	"fmt"
	"github.com/QWIDEX/order-proccesing/common"
	pb "github.com/QWIDEX/order-proccesing/common/api"
	_ "github.com/joho/godotenv/autoload"
	"net/http"
	"time"
)

var (
	Addr = common.GetEnv("PORT", "8080")
)

func NewServer(orderClient pb.OrderServiceClient) *http.Server {
	handler := NewHandler(orderClient)

	return &http.Server{
		Addr:         fmt.Sprintf(":%s", Addr),
		Handler:      handler.registerRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}
