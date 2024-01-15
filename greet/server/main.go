package main

import (
	"log"
	"net"

	pb "github.com/hirotofuu/grpc-experience/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

// GreetServiceのサーバーオブジェクト
type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	// grpcサーバー作成
	s := grpc.NewServer()

	// GreetServiceのサーバーをgrpcサーバーに接続
	pb.RegisterGreetServiceServer(s, &Server{})

	// grpcサーバーと開けたポートを繋げる
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
