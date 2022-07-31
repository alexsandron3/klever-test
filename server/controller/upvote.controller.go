package controller

import (
	"context"
	"fmt"
	"log"

	upvote "github.com/alexsandron3/klever-test/proto"
	"github.com/alexsandron3/klever-test/server/service"
)

// TO-DO = Change UnimplementedUpvoteServiceServer to type that have more semantic value
type Server struct {
	upvote.UnimplementedUpvoteServiceServer
}

func (s *Server) GetAllUsers(input *upvote.GetAllRequest, stream upvote.UpvoteService_GetAllUsersServer) error {
	allUsers, err := service.GetAll()

	fmt.Println(allUsers)
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range allUsers {

		res := &upvote.GetAllResponse{
			Name:  user.Name,
			Votes: user.Votes,
		}
		stream.Send(res)
	}
	return nil
}

func (s *Server) NewVote(ctx context.Context, input *upvote.NewVoteRequest) (*upvote.NewVoteResponse, error) {
	err := service.CheckIfIdIsValid(input.GetId())

	if err != nil {
		return nil, err
	}

	user, err := service.NewVote(input.GetId(), input.GetUpVote())

	if err != nil {
		return nil, err
	}

	return &upvote.NewVoteResponse{Name: user.Name, Votes: user.Votes}, nil
}
