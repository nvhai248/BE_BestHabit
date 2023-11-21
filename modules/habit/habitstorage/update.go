package habitstorage

import (
	"bestHabit/common"
	"bestHabit/modules/habit/habitmodel"
	"context"
)

func (s *sqlStore) UpdateHabitInfo(ctx context.Context, newInfo *habitmodel.HabitUpdate, id int) error {
	db := s.db

	startDate, err := common.ParseStringToDate(*newInfo.StartDate)
	if err != nil {
		return common.ErrInternal(err)
	}

	endDate, err := common.ParseStringToDate(*newInfo.EndDate)
	if err != nil {
		return common.ErrInternal(err)
	}

	reminderTime, err := common.ParseStringToTime(*newInfo.Reminder)
	if err != nil {
		return common.ErrInternal(err)
	}

	if _, err := db.Exec("UPDATE habits SET name = ?, description = ?, type = ?, start_date = ?, end_date = ?, reminder = ?, days = ? WHERE id = ?",
		newInfo.Name, newInfo.Description, newInfo.Type, startDate, endDate, reminderTime, newInfo.Days, id); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
