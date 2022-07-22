package main

import (
	"context"
	pb "github.com/rexposadas/learning-grpc/greet/proto"
	"io"
	"log"
	"time"
)

func doGreetEveryone(c pb.GreetServiceClient) {

	log.Println("invoked doGreetEveryone client")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Printf("failed to create client stream %v", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "jean"},
		{FirstName: "maria"},
		{FirstName: "last"},
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

			log.Printf("Received : %v\n", res.Result)
		}

		close(waitc)
	}()
	
	<-waitc
}
