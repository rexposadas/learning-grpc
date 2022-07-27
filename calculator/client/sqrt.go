package main

import (
	"context"
	pb "github.com/rexposadas/learning-grpc/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {

	log.Println("doSqrt invoked")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{
		Num: n,
	})

	if err != nil {
		e, ok := status.FromError(err)
		if ok { // grpc error
			log.Printf("error message from server: %s\n", e.Message())
			log.Printf("error code from server :%s", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Println("probably sent negative numbers")
				return
			}

		} else {
			log.Printf("no grpc error %v", e)
		}

	}

	log.Printf("sqrt : %f", res.Result)

}
