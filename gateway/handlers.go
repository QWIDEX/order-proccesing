package main

import (
	"github.com/QWIDEX/order-proccesing/common"
	pb "github.com/QWIDEX/order-proccesing/common/api"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
)

type handler struct {
	orderClient pb.OrderServiceClient
}

func NewHandler(orderClient pb.OrderServiceClient) *handler {
	return &handler{orderClient}
}

func (h *handler) registerRoutes() http.Handler {
	r := gin.Default()

	r.POST("/user/:uid/orders", h.handleCreateOrder)

	return r
}

func (h *handler) handleCreateOrder(c *gin.Context) {
	uid := c.Param("uid")

	var items []*pb.Item

	err := common.ReadJSON(c, &items)
	if err != nil {
		log.Println(err.Error())
		common.WriteError(c, common.ApiError{Code: http.StatusBadRequest, Message: "error while reading json"})
		return
	}

	_, err = h.orderClient.CreateOrder(c, &pb.CreateOrderReq{Uid: uid, Items: items})
	rStatus := status.Convert(err)

	if rStatus != nil {
		if rStatus.Code() != codes.InvalidArgument {
			common.WriteError(c, common.ApiError{Code: http.StatusBadRequest, Message: rStatus.Message()})
			return
		}

		log.Println(err)
		common.WriteError(c, common.ApiError{Code: http.StatusInternalServerError, Message: "something went wrong"})
		return
	}

	c.JSON(200, gin.H{})
}
