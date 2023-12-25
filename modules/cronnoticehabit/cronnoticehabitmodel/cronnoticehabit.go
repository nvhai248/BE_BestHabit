package cronnoticehabitmodel

import "github.com/robfig/cron/v3"

type CronNoticeHabit struct {
	UserId  int          `json:"user_id" db:"user_id"`
	EntryId cron.EntryID `json:"entry_id" db:"entry_id"`
	HabitId int          `json:"habit_id" db:"habit_id"`
}

func (CronNoticeHabit) TableName() string {
	return "cron_notice_habit"
}
