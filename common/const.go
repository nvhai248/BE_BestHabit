package common

const (
	DbTypeUser      = 1
	DbTypeHabit     = 2
	DbTypeTask      = 3
	DbTypeChallenge = 4
	DbParticipant   = 5
)

const CurrentUser = "user"

type Requester interface {
	GetId() int
	GetEmail() string
	GetRole() string
}
