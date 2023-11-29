package challengemodel

type ChallengeUpdate struct {
	Name            *string `json:"name" db:"name"`
	Description     *string `json:"description"`
	StartDate       *string `json:"start_date" db:"start_date"`
	EndDate         *string `json:"end_date" db:"end_date"`
	ExperiencePoint *int    `json:"experience_point" db:"experience_point"`
}

func (ChallengeUpdate) TableName() string {
	return Challenge{}.TableName()
}
