package controller

import (
	"io"
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

	if err != nil {
		log.Fatal(err)
	}
	for _, user := range allUsers {

		res := &upvote.GetAllResponse{
			Name:  user.Name,
			Votes: user.Votes,
			Id:    user.ID.Hex(),
		}
		stream.Send(res)
	}
	return nil
}

func (s *Server) NewVote(stream upvote.UpvoteService_NewVoteServer) error {

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}
		err = service.CheckIfIdIsValid(req.GetId())

		if err != nil {
			return err
		}

		user, err := service.NewVote(req.GetId(), req.GetUpVote())

		if err != nil {
			return err
		}
		stream.Send(&upvote.NewVoteResponse{Name: user.Name, Votes: user.Votes})

	}
}
