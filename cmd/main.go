package main

import (
	"grpc-crud-app/internal/repository"
	"grpc-crud-app/internal/server"
	"grpc-crud-app/internal/service"
	"log"
)

func main() {
	repo := repository.NewTodoRepository()
	todoService := service.NewTodoService(repo)

	err := server.StartGRPCServer(todoService, "8080")
	if err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}
}
