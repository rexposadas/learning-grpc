package main

import (
	"fmt"
	pb "github.com/rexposadas/learning-grpc/greet/proto"
	"log"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {

	log.Printf("greet many times invoked with %v", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("hello %s, number %d", in.FirstName, i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
