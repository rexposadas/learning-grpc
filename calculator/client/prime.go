package main

import (
	pb "github.com/rexposadas/learning-grpc/calculator/proto"
	"io"
	"log"

	"context"
)

func doPrime(c pb.CalculatorServiceClient) {
	log.Println("invoked prime client")

	req := &pb.PrimeRequest{
		Num: 120,
	}
	stream, err := c.Prime(context.Background(), req)

	if err != nil {
		log.Fatalf("get stream error %v", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error reading stream %v", err)
		}

		log.Printf("prime result %d", msg.Result)

	}

}
