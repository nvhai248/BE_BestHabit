package participantmodel

import "bestHabit/common"

type ParticipantFind struct {
	common.SQLModel `json:",inline"`
	UserId          int    `json:"user_id" db:"user_id"`
	ChallengeId     int    `json:"challenge_id" db:"challenge_id"`
	Status          string `json:"status" db:"status"`
}

func (ParticipantFind) TableName() string {
	return Participant{}.TableName()
}

func (t *ParticipantFind) Mask(isAdminOrOwner bool) {
	t.GenUID(common.DbParticipant)
}
