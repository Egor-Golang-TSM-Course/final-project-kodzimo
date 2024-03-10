package main

import (
	"context"
	"log"
	"net"

	pb "final-project-kodzimo/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHashingServer
}

func (s *server) CheckHash(ctx context.Context, in *pb.HashRequest) (*pb.HashResponse, error) {
	// Implement your logic here
}

func (s *server) GetHash(ctx context.Context, in *pb.HashRequest) (*pb.HashResponse, error) {
	// Implement your logic here
}

func (s *server) CreateHash(ctx context.Context, in *pb.HashRequest) (*pb.HashResponse, error) {
	// Implement your logic here
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHashingServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

/*
Этот код создает gRPC сервер и регистрирует ваш Hashing Service на этом сервере.
Затем он начинает слушать входящие запросы на порту 50051.
*/
