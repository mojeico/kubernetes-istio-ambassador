package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"time"

	pb "grpc-client/proto"
)

func main() {

	app := fiber.New()

	app.Get("/api/v1/http", HttpRequest)
	app.Get("/api/v1/grpc", RunGRPC)

	// Start server on port 3000
	app.Listen(":3000")

}

func HttpRequest(c *fiber.Ctx) error {

	var err error
	var resp *http.Response

	EndUserHeader := c.GetReqHeaders()["End-User"]

	if EndUserHeader != "" {
		client := &http.Client{}
		req, _ := http.NewRequest("GET", "http://http-server-service:3001/server/http", nil)
		req.Header.Set("End-User", EndUserHeader)
		resp, err = client.Do(req)
	} else {
		client := &http.Client{}
		req, _ := http.NewRequest("GET", "http://http-server-service:3001/server/http", nil)
		resp, err = client.Do(req)
	}

	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	var responseBody = ""

	scanner := bufio.NewScanner(resp.Body)

	for i := 0; scanner.Scan() && i < 5; i++ {
		responseBody = scanner.Text()
	}

	if resp.StatusCode == 500 {
		log.Println("-------")
		log.Println(resp.StatusCode)
		log.Println(responseBody)
		log.Println("-------")
	} else {
		log.Println(responseBody)
		log.Println(resp.StatusCode)
	}

	return c.JSON(fiber.Map{"error": responseBody})
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	todos := []TodoTask{
		{Name: "Meet with mentor", Description: "Discuss blockers in my project", Done: false},
	}

	for _, todo := range todos {
		_, err := c.CreateTodo(ctx, &pb.NewTodo{Name: todo.Name, Description: todo.Description, Done: todo.Done})

		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}

		log.Printf("messae was sent")
	}

	return nil

}
