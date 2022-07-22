package main

import (
	pb "github.com/rexposadas/learning-grpc/calculator/proto"
	"io"
	"log"
	"math"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {

	log.Println("max server invoked")

	highest := int32(math.MinInt32)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("error receiving %v", err)
		}

		if req.Num > highest {
			highest = req.Num
		}

		stream.Send(&pb.MaxResponse{
			Num: highest,
		})

	}
}
