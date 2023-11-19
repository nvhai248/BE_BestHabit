package taskmodel

import "bestHabit/common"

type TaskCreate struct {
	common.SQLModel `json:",inline"`
	UserId          int    `json:"user_id" db:"user_id"`
	Name            string `json:"name" db:"name"`
	Description     string `json:"description" db:"description"`
	Deadline        string `json:"deadline" db:"deadline"`
	Reminder        string `json:"reminder" db:"reminder"`
	Status          string `json:"status" db:"status"`
}

func (TaskCreate) TableName() string {
	return "tasks"
}

func (t *TaskCreate) Validate() error {
	if t.Name == "" {
		return ErrNameNotBeBlank
	}

	if t.Deadline == "" {
		return ErrDeadlineNotBeBlank
	}

	return nil
}
