package cronnoticetaskmodel

import "github.com/robfig/cron/v3"

type CronNoticeTask struct {
	UserId  int          `json:"user_id" db:"user_id"`
	EntryId cron.EntryID `json:"entry_id" db:"entry_id"`
	TaskId  int          `json:"task_id" db:"task_id"`
}

func (CronNoticeTask) TableName() string {
	return "cron_notice_task"
}
