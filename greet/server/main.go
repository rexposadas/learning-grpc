package main

import (
	"log"
	"net"

	pb "github.com/rexposadas/learning-grpc/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on  %v", err)
	}

	log.Println("listening on ", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Printf("failed to serve %v", err)
	}
}
