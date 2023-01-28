package main

import (
	"context"
	"log"
	"net"

	"github.com/google/uuid"
	"google.golang.org/grpc"

	pb "get-server-grpc/proto"
)

const (
	PORT = ":50052"
)

type TodoServer struct {
	pb.UnimplementedTodoServiceServer
}

func (s *TodoServer) CreateTodo(ctx context.Context, in *pb.NewTodo) (*pb.Todo, error) {
	log.Printf("--- v1 ---")
	todo := &pb.Todo{
		Name:        in.GetName(),
		Description: in.GetDescription(),
		Done:        false,
		Id:          uuid.New().String(),
	}

	return todo, nil

}

func main() {
	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("failed connection: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterTodoServiceServer(s, &TodoServer{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
