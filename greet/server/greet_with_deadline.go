package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"

	pb "github.com/rexposadas/learning-grpc/greet/proto"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {

	log.Printf("greet with deadline invoked with : %v", in)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("the client cancelled the request")
			return nil, status.Errorf(codes.Canceled, "client cancelled request")
		}

		time.Sleep(time.Second)
	}

	return &pb.GreetResponse{
		Result: " Hello " + in.FirstName,
	}, nil
}
