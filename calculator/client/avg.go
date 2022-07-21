package main

import (
	"context"
	pb "github.com/rexposadas/learning-grpc/calculator/proto"
	"log"
	"time"
)

func doAvg(c pb.CalculatorServiceClient) {

	log.Println("doAvg invoked")

	reqs := []*pb.AvgRequest{
		{Num: 10},
		{Num: 5},
		{Num: 100},
		{Num: 50},
	}

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("error while calling long greet %v", err)
	}

	for _, req := range reqs {
		log.Printf("sending req : %v\n", req)
		stream.Send(req)
		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Printf("error while receiving response from long greet %s", err)
	}

	log.Printf("ave : %f", res.Result)

}
