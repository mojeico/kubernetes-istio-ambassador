package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
	pb "grpc-client/proto"
	"log"
	"net/http"
	"time"
)

const (
	ADDRESS = "server-service:50051"
)

type TodoTask struct {
	Name        string
	Description string
	Done        bool
}

func main() {

	err := profiler.Start(
		profiler.WithService("golang-client"),
		profiler.WithEnv("snd"),
		profiler.WithVersion("v1"),
		profiler.WithTags("app:golang-client", "app1:golang-client1"),
		profiler.WithProfileTypes(
			profiler.CPUProfile,
			profiler.HeapProfile,
			// The profiles below are disabled by default to keep overhead
			// low, but can be enabled as needed.

			profiler.BlockProfile,
			profiler.MutexProfile,
			profiler.GoroutineProfile,
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer profiler.Stop()

	router := mux.NewRouter()

	router.Path("/metrics").Handler(promhttp.Handler())

	router.HandleFunc("/api/v1/http", HttpRequest)

	router.HandleFunc("/api/v1/grpc", RunGRPC)
	fmt.Println("Serving requests on port 3000")
	err = http.ListenAndServe(":3000", router)
	log.Fatal(err)
}

func HttpRequest(w http.ResponseWriter, r *http.Request) {

	var resp *http.Response

	//EndUserHeader := r.GetReqHeaders()["End-User"]

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://http-server-service:80/server/http", nil)
	resp, _ = client.Do(req)

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

	//return c.JSON(fiber.Map{"error": responseBody})

}

func RunGRPC(w http.ResponseWriter, r *http.Request) {
	//trace, _ := opentracing.StartSpanFromContext(r.Context(), "Handle /api/v1/grpc")
	//time.Sleep(time.Second / 2)
	//defer trace.Finish()

	//requestResultCounter.WithLabelValues("server", "200").Inc()

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

}
