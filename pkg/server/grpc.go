package server

import (
	"fmt"
	"net"

	"github.com/bag-huyag/user-service/internal/handler"
	"github.com/bag-huyag/user-service/internal/kafka"
	pb "github.com/bag-huyag/user-service/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGRPC() error {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	// producer := kafka.NewProducer("localhost:9092", "users")
	producer := kafka.NewProducer("kafka:9092", "users")

	srv := grpc.NewServer()
	reflection.Register(srv)
	pb.RegisterUserServiceServer(srv, handler.NewUserHandler(producer))

	fmt.Println("gRPC server listening on port 50052")
	return srv.Serve(lis)
}
