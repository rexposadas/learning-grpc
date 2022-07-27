package main

import (
	"context"
	pb "github.com/rexposadas/learning-grpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "clement",
	}

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			// related to grpc

			if e.Code() == codes.DeadlineExceeded {
				log.Println("deadline exceeded")
				return
			}

			log.Fatalf("unexpexted grpc error %v", err)
		} else {
			log.Printf("non grpc error %v", err)
		}
	}
	log.Printf("greet with deadline %s", res.Result)

}
