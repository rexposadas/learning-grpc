package main

import (
	"context"
	"log"

	pb "github.com/rexposadas/learning-grpc/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {

	log.Printf("greet func was invoked with %v", in)

	return &pb.GreetResponse{
		Result: "greetings, " + in.FirstName,
	}, nil
}
