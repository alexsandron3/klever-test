package service_test

import (
	"testing"

	"github.com/alexsandron3/klever-test/server/service"
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
