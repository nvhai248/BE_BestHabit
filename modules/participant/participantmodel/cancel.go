package participantmodel

type ParticipantCancel struct {
	UserId      int `json:"user_id" db:"user_id"`
	ChallengeId int `json:"challenge_id" db:"challenge_id"`
}

func (ParticipantCancel) TableName() string {
	return Participant{}.TableName()
}
