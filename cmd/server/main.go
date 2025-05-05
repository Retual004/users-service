package main

import (
	"log"

	database "github.com/Retual004/users-service/internal/database"
	transportgrpc "github.com/Retual004/users-service/internal/transport/grpc"
	"github.com/Retual004/users-service/internal/user"
)

func main() {
	database.InitDB()
	repo := user.NewUserRepository(database.DB)
	svc  := user.NewUserService(repo)

	if err := transportgrpc.RunGRPC(svc); err != nil {
	  log.Fatalf("gRPC server error: %v", err)
	}
}