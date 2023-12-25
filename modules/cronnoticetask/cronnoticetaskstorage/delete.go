package cronnoticetaskstorage

import (
	"bestHabit/common"
	"context"
)

func (s *sqlStore) DeleteCronNoticeTasks(ctx context.Context, userId, taskId int) error {
	db := s.db

	if _, err := db.Exec("DELETE FROM cron_notice_tasks WHERE user_id = ? AND task_id = ?)",
		userId, taskId); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
