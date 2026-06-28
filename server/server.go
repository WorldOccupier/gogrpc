package main

import (
	pb "climberinfo/protobuf"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var port = 8080

type server struct {
	pb.UnimplementedClimberInfoServer
}

func (s *server) GetClimber(_ context.Context, request *pb.GetClimberRequest) (*pb.GetClimberResponse, error) {
	return &pb.GetClimberResponse{Info: request.Name + " is great"}, nil
}

func run() {
	log.Println("Running grpc server")
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen on: %d", port)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterClimberInfoServer(grpcServer, &server{})
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("Failed to serve on: %d", port)
	}
}
