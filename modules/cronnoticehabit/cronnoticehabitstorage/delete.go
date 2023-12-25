package cronnoticehabitstorage

import (
	"bestHabit/common"
	"context"
)

func (s *sqlStore) DeleteCronNoticeHabits(ctx context.Context, userId, habitId int) error {
	db := s.db

	if _, err := db.Exec("DELETE FROM cron_notice_habits WHERE user_id = ? AND habit_id = ?)",
		userId, habitId); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
