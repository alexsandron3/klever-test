package model

import (
	"context"
	"encoding/json"
	"fmt"
	_ "fmt"
	_ "io/ioutil"
	_ "log"
	_ "net/http"

	upvote "github.com/alexsandron3/klever-test/proto"
	"github.com/alexsandron3/klever-test/server/controller"
	"go.mongodb.org/mongo-driver/bson"
)

// I was getting an problem when trying to start server, it always returned me that model.Server do not implement upvote.UpvoteServiceServer
type Server struct {
	upvote.UnimplementedUpvoteServiceServer
}

type User struct {
	ID    string `json:"_id"`
	Name  string `json:"name"`
	Votes int64  `json:"votes"`
}

// TO-DO = Refact this code to return all users
func (s *Server) GetAllUsers(ctx context.Context, in *upvote.GetAllRequest) (*upvote.GetAllResponse, error) {
	result := controller.GetAllUsers()
	jsonData, err := json.MarshalIndent(result, "", "    ")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
	return &upvote.GetAllResponse{
		Id:    "62e53702bb3b492963728230",
		Name:  "alexsandro",
		Votes: 000,
	}, nil
}

func (s *Server) NewVote(ctx context.Context, in *upvote.NewVoteService) (*upvote.NewVoteResponse, error) {

	result := controller.UpvoteUser(in.GetId())
	var user User
	bsonBytes, _ := bson.Marshal(result)
	bson.Unmarshal(bsonBytes, &user)

	return &upvote.NewVoteResponse{
		Name:  user.Name,
		Votes: user.Votes,
	}, nil
}
