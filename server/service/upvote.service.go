package service

import (
	"github.com/alexsandron3/klever-test/server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User struct {
	ID    string `json:"_id"`
	Name  string `json:"name"`
	Votes int64  `json:"votes"`
}

func GetVoteValue(upvote bool) int16 {
	voteValue := -1
	if upvote == true {
		voteValue = 1
	}
	return int16(voteValue)
}

func CheckIfIdIsValid(id string) error {
	if len(id) != 24 {
		return status.Errorf(codes.InvalidArgument, "Invalid ID")
	}
	return nil
}

func NewVote(userId string, vote bool) (user User, err error) {
	objId, err := primitive.ObjectIDFromHex(userId)
	result, err := model.UpdateVote(GetVoteValue(vote), objId)
	if err != nil {
		return user, err
	}
	bsonResult, _ := bson.Marshal(result)
	bson.Unmarshal(bsonResult, &user)

	return user, nil

}
