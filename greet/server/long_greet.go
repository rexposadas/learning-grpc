package main

import (
	"fmt"
	pb "github.com/rexposadas/learning-grpc/greet/proto"
	"io"
	"log"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("long greet was invoked")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("error while reading client stream %s", err)
		}

		log.Printf("receiving : %v", req)
		res += fmt.Sprintf("Hello %s!\n", req.FirstName)

	}

	return nil
}
