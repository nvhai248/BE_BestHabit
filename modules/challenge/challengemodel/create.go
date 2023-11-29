package challengemodel

type ChallengeCreate struct {
	Name            string `json:"name" db:"name"`
	Description     string `json:"description"`
	StartDate       string `json:"start_date" db:"start_date"`
	EndDate         string `json:"end_date" db:"end_date"`
	ExperiencePoint int    `json:"experience_point" db:"experience_point"`
}

func (ChallengeCreate) TableName() string {
	return Challenge{}.TableName()
}

func (t *ChallengeCreate) Validate() error {
	if t.Name == "" {
		return ErrNameNotBeBlank
	}

	if t.StartDate == "" {
		return ErrStartDateNotBlank
	}

	if t.EndDate == "" {
		return ErrEndDateNotBlank
	}

	return nil
}
