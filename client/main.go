package main

import (
	pb "climberinfo/protobuf"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var clientTarget = "localhost:8080"

func main() {
	connection, err := grpc.NewClient(clientTarget, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to %v", err)
	}

	defer connection.Close()

	client := pb.NewClimberInfoClient(connection)
	request := pb.GetClimberRequest{
		Name: "Adam Ondra",
	}
	reponse, err := client.GetClimber(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error requesting climber info: %v", err)
	}

	log.Println(reponse.Info)
}
