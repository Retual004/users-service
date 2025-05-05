// internal/transport/grpc/server.go
package grpc

import (
	"fmt"
	"net"

	userpb "github.com/Retual004/project-protos/proto/user"
	"github.com/Retual004/users-service/internal/user"
	"google.golang.org/grpc"
)

// RunGRPC поднимает gRPC-сервер на порту 50051.
func RunGRPC(svc *user.UserService) error {
	// Слушаем порт
	addr := ":50051"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", addr, err)
	}

	// Создаём сервер
	grpcServer := grpc.NewServer()

	// Регистрируем хендлер
	handler := NewHandler(svc)
	userpb.RegisterUserServiceServer(grpcServer, handler)

	fmt.Printf("gRPC server listening on %s\n", addr)
	// Старт
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve gRPC: %w", err)
	}
	return nil
}
