package main

import (
	"google.golang.org/grpc/credentials"
	"log"
	"time"

	pb "github.com/rexposadas/learning-grpc/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "localhost:50051"

func main() {
	tls := true

	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("error loading ca trust certification: %v", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("failed to connect %v", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doGreetWithDeadline(c, 5*time.Second)
}
