package main

import (
	"context"
	pb "github.com/rexposadas/learning-grpc/calculator/proto"
	"io"
	"log"
	"time"
)

func doMax(c pb.CalculatorServiceClient) {

	log.Println("invoked doMax")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Printf("failed to create client stream %v", err)
	}

	reqs := []*pb.MaxRequest{
		{Num: 1},
		{Num: 6},
		{Num: 10},
		{Num: 9},
		{Num: 11},
		{Num: 3},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("sending %v", req)
			stream.Send(req)
			time.Sleep(time.Second)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("failed to receive %v", err)
			}

			log.Printf("current max: %v", res.Num)
		}
	}()

	<-waitc

}
