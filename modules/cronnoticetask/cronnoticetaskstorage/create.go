package cronnoticetaskstorage

import (
	"bestHabit/common"
	"bestHabit/modules/cronnoticetask/cronnoticetaskmodel"
	"context"
)

func (s *sqlStore) CreateNewCronNoticeTask(ctx context.Context, data *cronnoticetaskmodel.CronNoticeTask) error {
	db := s.db

	if _, err := db.Exec("INSERT INTO cron_notice_tasks (user_id, entry_id, task_id) VALUES (?, ?, ?)",
		data.UserId, data.EntryId, data.TaskId); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
