package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
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

//func init() {
//prometheus.Register(totalRequests)
//prometheus.Register(responseStatus)
//prometheus.Register(httpDuration)
//prometheus.Register(requestResultCounter)
//}

func main() {
	router := mux.NewRouter()
	//router.Use(prometheusMiddleware)

	// Prometheus endpoint

	//http.Handle("/metrics", promhttp.Handler())
	//http.Handle("/prometheus", promhttp.Handler())

	//router.Path("/prometheus").Handler(promhttp.Handler())
	router.Path("/metrics").Handler(promhttp.Handler())

	router.HandleFunc("/api/v1/grpc", RunGRPC)
	fmt.Println("Serving requests on port 3000")
	err := http.ListenAndServe(":3000", router)
	log.Fatal(err)
}

func RunGRPC(w http.ResponseWriter, r *http.Request) {
	trace, _ := opentracing.StartSpanFromContext(r.Context(), "Handle /api/v1/grpc")
	//time.Sleep(time.Second / 2)
	defer trace.Finish()

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
