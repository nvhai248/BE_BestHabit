package participantmodel

type ParticipantCreate struct {
	UserId      int `json:"user_id" db:"user_id"`
	ChallengeId int `json:"challenge_id" db:"challenge_id"`
}

func (ParticipantCreate) TableName() string {
	return Participant{}.TableName()
}
