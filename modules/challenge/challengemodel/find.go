package challengemodel

import "bestHabit/common"

type ChallengeFind struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" db:"name"`
	Description     string `json:"description"`
	StartDate       string `json:"start_date" db:"start_date"`
	EndDate         string `json:"end_date" db:"end_date"`
	ExperiencePoint int    `json:"experience_point" db:"experience_point"`
	Status          bool   `json:"status" db:"status"`
	CountUserJoined int    `json:"count_user_joined" db:"count_user_joined"`
}

func (ChallengeFind) TableName() string {
	return Challenge{}.TableName()
}

func (t *ChallengeFind) Mask(isAdminOrOwner bool) {
	t.GenUID(common.DbTypeChallenge)
}
