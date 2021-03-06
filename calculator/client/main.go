package main

import (
	"log"

	pb "github.com/rexposadas/learning-grpc/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to connect %v", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	doSqrt(c, -10)
}
