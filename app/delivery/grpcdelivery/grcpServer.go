package grpcdelivery

import (
	"context"
	"fmt"
	"modcleanarch/app/application/service"
	"modcleanarch/app/domain/entity"
	"modcleanarch/app/domain/usecase"
	"modcleanarch/app/infrastructure/database"
	databaseadapter "modcleanarch/app/infrastructure/databaseAdapter"
	"modcleanarch/grpc-config/pb"
)

type GrpcServer struct {
	PedidoUseCase service.IProdutoService
	pb.UnimplementedPedidoServiceServer
}

func (s *GrpcServer) ListarPedidos(ctx context.Context, req *pb.ListarPedidosRequest) (*pb.ListarPedidosResponse, error) {
	//Dependências
	s.PedidoUseCase = &usecase.ProdutoUseCase{
		PedidoRepository: &database.PedidoRepository{
			Conn: &databaseadapter.MariaDbConectar{},
		},
	}

	//Consulta os pedidos
	pedidos, err := s.PedidoUseCase.ListarPedidos()
	if err != nil {
		fmt.Println("Falha ao consultar pedidos -> REST:")
		fmt.Println(err)
		println()
	}

	result := s.convertToPbPedido(pedidos)

	return &pb.ListarPedidosResponse{Pedidos: result}, nil
}

func (s *GrpcServer) convertToPbPedido(pedidos []entity.Pedido) []*pb.Pedido {
	var pbPedidos []*pb.Pedido

	for _, pedido := range pedidos {
		pbPedido := &pb.Pedido{
			Id:            int32(pedido.ID),
			NumeroPedido:  int32(pedido.NumeroPedido),
			NomeProduto:   pedido.NomeProduto,
			Quantidade:    int32(pedido.Quantidade),
			PrecoUnitario: float32(pedido.PrecoUnitario),
		}
		pbPedidos = append(pbPedidos, pbPedido)
	}

	return pbPedidos
}
