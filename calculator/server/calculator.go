package main

import (
	"context"
	"log"

	pb "github.com/rexposadas/learning-grpc/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {

	log.Printf("greet func was invoked with %v", in)

	var sum int32
	for _, n := range in.Num {
		sum += n
	}

	return &pb.SumResponse{
		Result: sum,
	}, nil
}
