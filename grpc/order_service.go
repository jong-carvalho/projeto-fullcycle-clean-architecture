package grpc

//
//import (
//	"context"
//	"projeto-fullcycle-clean-architecture/proto"    // Importa o pacote protobuf gerado
//	"projeto-fullcycle-clean-architecture/services" // Importa os serviços de aplicação
//)
//
//// Define o struct que implementa a interface gRPC gerada
//type OrderServiceServer struct {
//	proto.UnimplementedOrderServiceServer
//}
//
//// ListOrders é a implementação do método gRPC que lista os pedidos
//func (s *OrderServiceServer) ListOrders(ctx context.Context, req *proto.EmptyRequest) (*proto.OrderListResponse, error) {
//	// Chama o serviço para listar os pedidos
//	orders := services.ListOrders()
//
//	// Prepara a resposta para o formato esperado pelo gRPC
//	var response []*proto.OrderResponse
//	for _, order := range orders {
//		response = append(response, &proto.OrderResponse{
//			Id:           uint32(order.ID), // Converte o ID para uint32
//			CustomerName: order.CustomerName,
//			TotalAmount:  order.TotalAmount,
//		})
//	}
//
//	// Retorna a resposta do gRPC
//	return &proto.OrderListResponse{Orders: response}, nil
//}
