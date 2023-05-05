package grcp

import (
	pb "apigateway/protos"
	"log"

	"google.golang.org/grpc"
)

func Setup() *grpc.ClientConn {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return conn
}

func NewAuthServiceClient(conn *grpc.ClientConn) pb.AuthServiceClient {
	return pb.NewAuthServiceClient(conn)
}
