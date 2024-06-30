package main

import (
	"log"
	"net"

	infraGrpc "github.com/myusername/myservice/internal/infrastructure/grpc"
	"github.com/myusername/myservice/internal/interface/repository"
	"github.com/myusername/myservice/internal/usecase"
	"github.com/myusername/myservice/proto"

	"google.golang.org/grpc"
)

func main() {
	// リポジトリを初期化
	userRepo := repository.NewUserRepository()

	// ユースケースを初期化
	userUsecase := usecase.NewUserUsecase(userRepo)

	// gRPCハンドラーを初期化
	grpcHandler := infraGrpc.NewUserHandler(userUsecase)

	// gRPCサーバーを作成
	server := grpc.NewServer()
	proto.RegisterUserServiceServer(server, grpcHandler)

	// TCPリスナーを作成
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Server is running on port :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
