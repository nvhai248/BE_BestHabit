package subscriber

type HasUserId interface {
	GetUserId() int
}

type HasChallengeId interface {
	GetChallengeId() int
}
