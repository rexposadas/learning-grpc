package main

import (
	"context"
	pb "github.com/rexposadas/learning-grpc/greet/proto"
	"log"
	"time"
)

func doLongGreet(c pb.GreetServiceClient) {

	log.Println("do long greet invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "jean"},
		{FirstName: "marie"},
		{FirstName: "last-one"},
	}

	stream, err := c.LongGreet(context.Background())

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

	log.Printf("long greet : %s", res.Result)

}
