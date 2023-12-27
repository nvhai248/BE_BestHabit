package habitmodel

import "bestHabit/common"

type HabitUpdate struct {
	Name           *string               `json:"name" db:"name"`
	Description    *string               `json:"description" db:"description"`
	StartDate      *string               `json:"start_date" db:"start_date"`
	EndDate        *string               `json:"end_date" db:"end_date"`
	Type           *string               `json:"type" db:"type"`
	Reminder       *string               `json:"reminder" db:"reminder"`
	IsCountBased   *bool                 `json:"is_count_based" db:"is_count_based"`
	CompletedDates *common.CompleteDates `json:"completed_dates" db:"completed_dates"`
	Days           *common.Days          `json:"days" db:"days"`
	Target         *common.Target        `json:"target" db:"target"`
	UserId         *int
	Id             *int
}

func (HabitUpdate) TableName() string {
	return Habit{}.TableName()
}

func (t *HabitUpdate) GetUserId() int {
	return *t.UserId
}

func (t *HabitUpdate) GetDescription() string {
	return *t.Description
}

func (t *HabitUpdate) GetName() string {
	return *t.Name
}

func (t *HabitUpdate) GetReminderTime() string {
	return *t.Reminder
}

func (t *HabitUpdate) GetStartDate() string {
	return *t.StartDate
}

func (t *HabitUpdate) GetEndDate() string {
	return *t.EndDate
}

func (t *HabitUpdate) GetDays() *common.Days {
	return t.Days
}

func (t *HabitUpdate) GetHabitId() int {
	return *t.Id
}
