package habitmodel

import "bestHabit/common"

type HabitFind struct {
	common.SQLModel `json:",inline"`
	UserId          int            `json:"-" db:"user_id"`
	Name            string         `json:"name" db:"name"`
	Description     string         `json:"description" db:"description"`
	StartDate       string         `json:"start_date" db:"start_date"`
	EndDate         string         `json:"end_date" db:"end_date"`
	Type            string         `json:"type" db:"type"`
	Reminder        string         `json:"reminder" db:"reminder"`
	Status          int            `json:"status" db:"status"`
	IsCountBased    bool           `json:"is_count_based" db:"is_count_based"`
	CompletedDates  *common.Dates  `json:"completed_dates" db:"completed_dates"`
	Days            *common.Days   `json:"days" db:"days"`
	Target          *common.Target `json:"target" db:"target"`
}

func (HabitFind) TableName() string {
	return Habit{}.TableName()
}

func (t *HabitFind) Mask(isAdminOrOwner bool) {
	t.GenUID(common.DbTypeTask)
}
