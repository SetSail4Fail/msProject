package grpc

import (
	"context"
	"log"
	"msProject/config"
	"msProject/internal/sql/account"
	gen "msProject/mypkg/grpc/grpcGenerated"
	"msProject/mypkg/postgres"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	gen.UnimplementedGreeterServer 
}

const ConfigPath = "config/config.yaml"

func (s *server) CreateAcc(ctx context.Context, in *gen.UserInputRequest) (*gen.Reply, error) {
	cfg, err := config.NewConfig(ConfigPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	postgres.DbConnect(cfg)

	req := account.CreateAccRequest{
        Name:     in.GetName(),
        Password: in.GetPassword(),
        Email:    in.GetEmail(),
    }
	account.CreateAcc(cfg, req)
	return &gen.Reply{Message: "Hello, " + in.GetName() + ". Your account successfully created."}, nil
}

func CreateTCP() {
	// Создаем TCP-листенер на порту 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Создаем gRPC-сервер
	s := grpc.NewServer()

	// Регистрируем наш сервис на сервере
	gen.RegisterGreeterServer(s, &server{})
	reflection.Register(s)

	// Запускаем сервер
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}