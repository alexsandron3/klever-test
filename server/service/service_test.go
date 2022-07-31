package service_test

import (
	"testing"

	"github.com/alexsandron3/klever-test/server/service"
)

func TestGetVoteValue(t *testing.T) {
	upvoteValues := []bool{false, true}
	expectedVoteValues := []int16{-1, 1}

	for index, upvoteValue := range upvoteValues {
		voteValue := service.GetVoteValue(upvoteValue)
		expectedVoteValue := expectedVoteValues[index]
		if voteValue != expectedVoteValue {
			t.Errorf("Expected: %d Received: %d ", expectedVoteValue, voteValue)
		}

	}
}
