package cronnoticehabitstorage

import (
	"bestHabit/common"
	"bestHabit/modules/cronnoticehabit/cronnoticehabitmodel"
	"context"
)

func (s *sqlStore) CreateNewCronNoticeHabit(ctx context.Context, data *cronnoticehabitmodel.CronNoticeHabit) error {
	db := s.db

	if _, err := db.Exec("INSERT INTO cron_notice_habits (user_id, entry_id, habit_id) VALUES (?, ?, ?)",
		data.UserId, data.EntryId, data.HabitId); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
