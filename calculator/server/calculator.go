package main

import (
	"context"
	"log"

	pb "github.com/rexposadas/learning-grpc/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {

	log.Printf("greet func was invoked with %v", in)

	return &pb.SumResponse{
		Result: in.Num + in.Num2,
	}, nil
}
