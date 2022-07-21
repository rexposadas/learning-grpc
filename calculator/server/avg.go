package main

import (
	pb "github.com/rexposadas/learning-grpc/calculator/proto"
	"io"
	"log"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {

	var sum int32
	var count int32

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: getAverage(sum, count),
			})
		}

		if err != nil {
			log.Fatalf("error while reading client stream %s", err)
		}

		count++
		sum += req.Num

		log.Printf("receiving : %v, sum: %d", req, sum)
	}

	return nil
}

func getAverage(total, count int32) float64 {

	if count == 0 {
		return 0
	}

	return float64(total / count)
}
