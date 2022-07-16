package main

import (
	"context"
	pb "github.com/rexposadas/learning-grpc/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "rex",
	})

	if err != nil {
		log.Fatalf("could not greet : %v", err)
	}

	log.Printf("Greeting: %s", res.Result)
}
