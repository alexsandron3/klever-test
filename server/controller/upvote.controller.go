package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	upvote "github.com/alexsandron3/klever-test/proto"
	"github.com/alexsandron3/klever-test/server/model"
	"github.com/alexsandron3/klever-test/server/service"
)

// TO-DO = Change UnimplementedUpvoteServiceServer to type that have more semantic value
type Server struct {
	upvote.UnimplementedUpvoteServiceServer
}

// TO-DO = Refact this code to return all users
func (s *Server) GetAllUsers(ctx context.Context, input *upvote.GetAllRequest) (*upvote.GetAllResponse, error) {
	allUsers := model.GetAllUsers()

	jsonData, err := json.Marshal(allUsers)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", jsonData)

	return &upvote.GetAllResponse{
		Id:    "62e53702bb3b492963728230",
		Name:  "alexsandro",
		Votes: 000,
	}, nil
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
