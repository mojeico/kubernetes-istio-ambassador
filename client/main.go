package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"

	pb "grpc-client/proto"
	"log"
	"net/http"
	"time"
)

type TodoTask struct {
	Name        string
	Description string
	Done        bool
}

func main() {

	router := mux.NewRouter()

	router.Path("/metrics").Handler(promhttp.Handler())

	router.HandleFunc("/api/v1/http", HttpRequest)

	router.HandleFunc("/api/v1/grpc", RunGRPC)

	router.HandleFunc("/api/v1/resp", GetOldResponse)

	router.HandleFunc("/api/v1/resp/new", GetNewResponse)

	router.HandleFunc("/api/v1/header/list/all", GetHeaderList)

	fmt.Println("Serving requests on port 3000")
	err := http.ListenAndServe(":3000", router)
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

	//ServerAddress := os.Getenv("SERVER_ADDRESS")
	ServerAddress := "apidev.np.digitalzenith.io"

	fmt.Println("start RunGRPC1")

	conn, err := grpc.Dial(ServerAddress, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}

	fmt.Println("start RunGRPC2")

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

func GetOldResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Old response ")
	w.WriteHeader(http.StatusOK)
}

func GetNewResponse(w http.ResponseWriter, r *http.Request) {

	json := simplejson.New()
	json.Set("foo", "bar")

	payload, err := json.MarshalJSON()

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func GetHeaderList(w http.ResponseWriter, r *http.Request) {

	headers := r.Header

	json := simplejson.New()

	for s, strings := range headers {
		json.Set(s, strings)
	}
	payload, err := json.MarshalJSON()

	if err != nil {
		log.Println(err)
	}

	w.Write(payload)

}
