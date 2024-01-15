package main

import (
	"context"
	"log"

	pb "github.com/hirotofuu/grpc-experience/greet/proto"
)

// GreetServiceのtype serverのメソッドに登録
func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet was invoked with: %v\n", in)

	return &pb.GreetResponse{
		Result: "Hello" + in.FirstName,
	}, nil
}
