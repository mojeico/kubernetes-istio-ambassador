package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	pb "grpc-server/proto"
)

const (
	PORT        = ":50051"
	GET_ADDRESS = "get-server-grpc:50052"
	//GET_ADDRESS = ":50052"
)

type TodoServer struct {
	pb.UnimplementedTodoServiceServer
}

type TodoTask struct {
	Name        string
	Description string
	Done        bool
}

func (s *TodoServer) CreateTodo(ctx context.Context, in *pb.NewTodo) (*pb.Todo, error) {
	log.Printf("Received: %v", in.GetName())
	todo := &pb.Todo{
		Name:        in.GetName(),
		Description: in.GetDescription(),
		Done:        false,
		Id:          uuid.New().String(),
	}

	log.Println("Hello world!")

	fmt.Println("start RunGRPC - CreateTodo")

	conn, err := grpc.Dial(GET_ADDRESS, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}

	defer conn.Close()

	c := pb.NewTodoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)

	defer cancel()

	todos := []TodoTask{
		{Name: "Code review", Description: "Review new feature code", Done: false},
	}

	for _, todo := range todos {
		res, err := c.CreateTodo(ctx, &pb.NewTodo{Name: todo.Name, Description: todo.Description, Done: todo.Done})

		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}

		log.Printf(`
           ID : %s
           Name : %s
           Description : %s
           Done : %v,
       `, res.GetId(), res.GetName(), res.GetDescription(), res.GetDone())
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
