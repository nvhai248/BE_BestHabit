package participantmodel

import "bestHabit/common"

const EntityName = "Participant"

type Participant struct {
	common.SQLModel `json:",inline"`
	UserId          int    `json:"user_id" db:"user_id"`
	ChallengeId     int    `json:"challenge_id" db:"challenge_id"`
	Status          string `json:"status" db:"status"`
}

func (Participant) TableName() string {
	return "participants"
}

func (t *Participant) Mask(isAdminOrOwner bool) {
	t.GenUID(common.DbParticipant)
}
