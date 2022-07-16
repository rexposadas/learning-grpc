package main

import (
	"context"
	pb "github.com/rexposadas/learning-grpc/calculator/proto"
	"log"
)

func doSum(c pb.SumServiceClient) {
	log.Println("doSum invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		Num: []int32{2, 345, 5},
	})

	if err != nil {
		log.Fatalf("could not sum : %v", err)
	}

	log.Printf("sum is: %d", res.Result)
}
