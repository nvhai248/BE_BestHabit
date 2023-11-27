package habitstorage

import (
	"bestHabit/common"
	"bestHabit/modules/habit/habitmodel"
	"context"
)

func (s *sqlStore) UpdateHabitInfo(ctx context.Context, newInfo *habitmodel.HabitUpdate, id int) error {
	db := s.db

	if _, err := db.Exec("UPDATE habits SET name = ?, description = ?, type = ?, start_date = ?, end_date = ?, reminder = ?, days = ? WHERE id = ?",
		newInfo.Name, newInfo.Description, newInfo.Type, newInfo.StartDate, newInfo.EndDate, newInfo.Reminder, newInfo.Days, id); err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) IncreaseCountCompleted(ctx context.Context, id int) error {
	db := s.db

	if _, err := db.Exec("UPDATE habits SET count_completed = count_completed + 1 WHERE id = ?",
		id); err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) UpdateCompletedDate(ctx context.Context, data *habitmodel.HabitUpdate, id int) error {
	db := s.db

	query := `UPDATE habits SET completed_dates = ? WHERE id = ?`

	if _, err := db.Exec(query,
		data.CompletedDates,
		id,
	); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
