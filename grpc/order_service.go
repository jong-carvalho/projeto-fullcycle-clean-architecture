package grpc

import (
	"context"
	pb "orders-service/proto"
	"orders-service/services"
)

type OrderServiceServer struct {
	pb.UnimplementedOrderServiceServer
}

func (s *OrderServiceServer) ListOrders(ctx context.Context, req *pb.EmptyRequest) (*pb.OrderListResponse, error) {
	orders := services.ListOrders()
	var response []*pb.OrderResponse

	for _, order := range orders {
		response = append(response, &pb.OrderResponse{
			Id:           uint32(order.ID),
			CustomerName: order.CustomerName,
			TotalAmount:  order.TotalAmount,
		})
	}

	return &pb.OrderListResponse{Orders: response}, nil
}
