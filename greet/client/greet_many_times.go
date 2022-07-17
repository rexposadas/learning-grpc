package main

import (
	"context"
	pb "github.com/rexposadas/learning-grpc/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("invoked doGreetManyTimes")

	req := &pb.GreetRequest{
		FirstName: "rex",
	}
	stream, err := c.GreetManyTimes(context.Background(), req)

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

		log.Printf("greet many times %s", msg.Result)

	}

}
