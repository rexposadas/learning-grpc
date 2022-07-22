package main

import (
	pb "github.com/rexposadas/learning-grpc/greet/proto"
	"io"
	"log"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {

	log.Println("invoked greet everyone server")

	for {

		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("error reading client steam %v", err)
		}

		res := "hello " + req.FirstName + "!"

		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("error sendin response %v", err)
		}
	}
}
