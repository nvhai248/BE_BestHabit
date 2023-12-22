package habitmodel

import "bestHabit/common"

type HabitCreate struct {
	UserId         int                   `json:"user_id" db:"user_id"`
	Name           string                `json:"name" db:"name"`
	Description    string                `json:"description" db:"description"`
	StartDate      string                `json:"start_date" db:"start_date"`
	EndDate        string                `json:"end_date" db:"end_date"`
	Type           string                `json:"type" db:"type"`
	Reminder       string                `json:"reminder" db:"reminder"`
	IsCountBased   bool                  `json:"is_count_based" db:"is_count_based"`
	Days           *common.Days          `json:"days" db:"days"`
	CompletedDates *common.CompleteDates `json:"completed_dates" db:"completed_dates"`
	Target         *common.Target        `json:"target" db:"target"`
}

func (HabitCreate) TableName() string {
	return Habit{}.TableName()
}

func (t *HabitCreate) GetUserId() int {
	return t.UserId
}

func (t *HabitCreate) Validate() error {
	if t.Name == "" {
		return ErrNameNotBeBlank
	}

	if t.StartDate == "" {
		return ErrStartDateNotBeBlank
	}

	if t.EndDate == "" {
		return ErrEndDateNotBeBlank
	}

	return nil
}
