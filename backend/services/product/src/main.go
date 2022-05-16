package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	pb "github.com/woriheck/go-commerce/shared/pricing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	http.HandleFunc("/product", HelloProduct)
	http.ListenAndServe(":8080", nil)
}

type Response struct {
	Message string `json:"message"`
}

func HelloProduct(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	msg := callRPC()

	jsonOut, _ := json.Marshal(Response{Message: msg})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", jsonOut)
}

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "grpc:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func callRPC() string {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return fmt.Sprintf("Greeting %s", r.GetMessage())
}
