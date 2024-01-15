package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/hirotofuu/grpc-experience/greet/proto"
)

var addr string = "localhost:50051"

func main() {

	// Dial creates a client connection to the given target.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	// rpcエンドポイントに接続できるクライアントを返す
	c := pb.NewGreetServiceClient(conn)

	doGreet(c)

}
