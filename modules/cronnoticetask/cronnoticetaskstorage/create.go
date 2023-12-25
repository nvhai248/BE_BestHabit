package cronnoticetaskstorage

import (
	"bestHabit/common"
	"bestHabit/modules/cronnoticetask/cronnoticetaskmodel"
	"context"

	"github.com/robfig/cron/v3"
)

func (s *sqlStore) CreateNewCronNoticeTask(ctx context.Context, data *cronnoticetaskmodel.CronNoticeTask) error {
	db := s.db

	if _, err := db.Exec("INSERT INTO cron_notice_tasks (user_id, entry_id, task_id) VALUES (?, ?, ?)",
		data.UserId, data.EntryId, data.TaskId); err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) CreateNewCronNoticeTasks(ctx context.Context, userId, taskId int, entryIds []cron.EntryID) error {
	for entryId := range entryIds {
		s.CreateNewCronNoticeTask(ctx, &cronnoticetaskmodel.CronNoticeTask{
			UserId:  userId,
			TaskId:  taskId,
			EntryId: cron.EntryID(entryId),
		})
	}
	return nil
}
