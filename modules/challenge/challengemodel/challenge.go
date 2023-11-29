package challengemodel

import "bestHabit/common"

const EntityName = "Challenge"

type Challenge struct {
	common.SQLModel `json:",inline"`
	Description     string `json:"description"`
	Name            string `json:"name" db:"name"`
	StartDate       string `json:"start_date" db:"start_date"`
	EndDate         string `json:"end_date" db:"end_date"`
	ExperiencePoint int    `json:"experience_point" db:"experience_point"`
	Status          bool   `json:"status" db:"status"`
}

func (Challenge) TableName() string {
	return "challenges"
}

func (t *Challenge) Mask(isAdminOrOwner bool) {
	t.GenUID(common.DbTypeChallenge)
}
