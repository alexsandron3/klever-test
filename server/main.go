package main

import (
	upvote "github.com/alexsandron3/klever-test/proto"
	model "github.com/alexsandron3/klever-test/server/model"

	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":8200")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	upv := model.Server{}

	grpcServer := grpc.NewServer()
	upvote.RegisterUpvoteServiceServer(grpcServer, &upv)

	log.Println("Listening on Port: 8200!")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
