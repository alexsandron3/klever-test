package service

func GetVoteValue(upvote bool) int16 {
	voteValue := -1
	if upvote == true {
		voteValue = 1
	}
	return int16(voteValue)
}
