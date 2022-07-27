package main

import (
	"context"
	"fmt"
	"log"
	"math"

	pb "github.com/rexposadas/learning-grpc/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(c context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("sqrt was invoked with %v", in)

	number := in.Num

	if number < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("received a negative number: %d", in.Num))
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}
