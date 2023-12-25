package cronnoticehabitstorage

import (
	"bestHabit/common"
	"bestHabit/modules/cronnoticehabit/cronnoticehabitmodel"
	"context"

	"github.com/robfig/cron/v3"
)

func (s *sqlStore) CreateNewCronNoticeHabit(ctx context.Context, data *cronnoticehabitmodel.CronNoticeHabit) error {
	db := s.db

	if _, err := db.Exec("INSERT INTO cron_notice_habits (user_id, entry_id, habit_id) VALUES (?, ?, ?)",
		data.UserId, data.EntryId, data.HabitId); err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) CreateNewCronNoticeHabits(ctx context.Context, userId, habitId int, entryIds []cron.EntryID) error {
	for entryId := range entryIds {
		s.CreateNewCronNoticeHabit(ctx, &cronnoticehabitmodel.CronNoticeHabit{
			UserId:  userId,
			HabitId: habitId,
			EntryId: cron.EntryID(entryId),
		})
	}
	return nil
}
