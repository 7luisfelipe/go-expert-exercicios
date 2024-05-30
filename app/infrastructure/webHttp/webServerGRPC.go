package webhttp

import (
	"log"
	"modcleanarch/app/delivery/grpcdelivery"
	"modcleanarch/grpc-config/pb"
	"net"

	"google.golang.org/grpc"
)

func WebGrpc() {
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("Falha - listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPedidoServiceServer(s, &grpcdelivery.GrpcServer{})
	log.Printf("Server gRPC rodando na porta: 8082")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Falha no server gRPC: %v", err)
	}
}
