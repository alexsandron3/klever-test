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
)

// I was getting an problem when trying to start server, it always returned me that model.Server do not implement upvote.UpvoteServiceServer
type Server struct {
	upvote.UnimplementedUpvoteServiceServer
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
