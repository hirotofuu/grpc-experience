package main

import (
	"context"
	"log"

	pb "github.com/hirotofuu/grpc-experience/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Hiroto",
	})

	if err != nil {
		log.Fatalf("could not greet: %v\n", err)
	}

	log.Printf("%s\n", res.Result)
}
