package cronnoticehabitstorage

import (
	"bestHabit/modules/cronnoticehabit/cronnoticehabitmodel"
	"context"
)

func (s *sqlStore) getListCronNoticeHabit(ctx context.Context, userId, habitId int) ([]cronnoticehabitmodel.CronNoticeHabit, error) {
	db := s.db

	var result []cronnoticehabitmodel.CronNoticeHabit

	if err := db.Select(&result, "SELECT * FROM cron_notice_habits WHERE user_id = ? AND habit_id = ?", userId, habitId); err != nil {
		return nil, err
	}

	return result, nil
}
