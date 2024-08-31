package server

import (
	"grpc-crud-app/api/proto/todo"
	"grpc-crud-app/internal/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func StartGRPCServer(service *service.TodoService, port string) error {

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	todo.RegisterTodoServiceServer(grpcServer, service)

	log.Printf("Starting gRPC server on port :%s", port)
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}
	return nil
}
