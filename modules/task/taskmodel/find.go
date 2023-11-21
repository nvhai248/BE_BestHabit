package taskmodel

import "bestHabit/common"

type TaskFind struct {
	common.SQLModel `json:",inline"`
	UserId          int    `json:"-" db:"user_id"`
	Name            string `json:"name" db:"name"`
	Description     string `json:"description" db:"description"`
	Deadline        string `json:"deadline" db:"deadline"`
	Reminder        string `json:"reminder" db:"reminder"`
	Status          string `json:"status" db:"status"`
}

func (TaskFind) TableName() string {
	return Task{}.TableName()
}

func (t *TaskFind) Mask(isAdminOrOwner bool) {
	t.GenUID(common.DbTypeTask)
}
