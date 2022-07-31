package service_test

import (
	"testing"

	"github.com/alexsandron3/klever-test/server/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetVoteValueWhenTrue(t *testing.T) {
	upvoteValue := true
	var expectedVoteValue int16 = 1

	voteValue := service.GetVoteValue(upvoteValue)
	if voteValue != expectedVoteValue {
		t.Errorf("Expected: %d Received: %d ", expectedVoteValue, voteValue)
	}
}

func TestGetVoteValueWhenFalse(t *testing.T) {
	upvoteValue := false
	var expectedVoteValue int16 = -1

	voteValue := service.GetVoteValue(upvoteValue)
	if voteValue != expectedVoteValue {
		t.Errorf("Expected: %d Received: %d ", expectedVoteValue, voteValue)
	}
}

func TestWhenIsIsInvalid(t *testing.T) {
	id := "1111111111111111111111111"

	err := service.CheckIfIdIsValid(id)

	if err == nil {
		t.Errorf("Expected to raise an error")
	}
}

func TestIfRaisedErrorIsCorrect(t *testing.T) {
	id := "1111111111111111111111111"

	err := service.CheckIfIdIsValid(id)
	expectedErr := status.Errorf(codes.InvalidArgument, "Invalid ID")
	// fmt.Println(err)
	if err.Error() != expectedErr.Error() {
		t.Errorf("Expected to raise correct error")
	}
}

func TestWhenIdIsValid(t *testing.T) {
	id := "111111111111111111111111"

	err := service.CheckIfIdIsValid(id)

	if err != nil {
		t.Errorf("Expected to not raise error")
	}
}
