package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"log"
	"time"

	pb "grpc-client/proto"
)

func main() {

	app := fiber.New()

	app.Get("/api/v1/post", GetPost)
	app.Get("/api/v1/grpc", RunGRPC)

	// Start server on port 3000
	app.Listen(":3000")

}

func GetPost(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"error": "Post not found"})
}

const (
	ADDRESS = "server-service:50051"
)

type TodoTask struct {
	Name        string
	Description string
	Done        bool
}

func RunGRPC(fiber *fiber.Ctx) error {

	log.Println("Hello world!")

	fmt.Print("start RunGRPC")

	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}

	defer conn.Close()

	c := pb.NewTodoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	todos := []TodoTask{
		{Name: "Code review", Description: "Review new feature code", Done: false},
		{Name: "Make YouTube Video", Description: "Start Go for beginners series", Done: false},
		{Name: "Go to the gym", Description: "Leg day", Done: false},
		{Name: "Buy groceries", Description: "Buy tomatoes, onions, mangos", Done: false},
		{Name: "Meet with mentor", Description: "Discuss blockers in my project", Done: false},
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

	return nil

}
