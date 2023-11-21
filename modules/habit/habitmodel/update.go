package habitmodel

import "bestHabit/common"

type HabitUpdate struct {
	Name        *string      `json:"name" db:"name"`
	Description *string      `json:"description" db:"description"`
	StartDate   *string      `json:"start_date" db:"start_date"`
	EndDate     *string      `json:"end_date" db:"end_date"`
	Type        *string      `json:"type" db:"type"`
	Reminder    *string      `json:"reminder" db:"reminder"`
	Days        *common.Days `json:"days" db:"days"`
}

func (HabitUpdate) TableName() string {
	return Habit{}.TableName()
}
