package cronnoticetaskstorage

import (
	"bestHabit/modules/cronnoticetask/cronnoticetaskmodel"
	"context"
)

func (s *sqlStore) getListCronNoticeTask(ctx context.Context, userId, taskId int) ([]cronnoticetaskmodel.CronNoticeTask, error) {
	db := s.db

	var result []cronnoticetaskmodel.CronNoticeTask

	if err := db.Select(&result, "SELECT * FROM cron_notice_tasks WHERE user_id = ? AND task_id = ?", userId, taskId); err != nil {
		return nil, err
	}

	return result, nil
}
