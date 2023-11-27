package habitstorage

import (
	"bestHabit/common"
	"bestHabit/modules/habit/habitmodel"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *habitmodel.HabitCreate) error {
	db := s.db

	query := `INSERT INTO habits (name, user_id, description, start_date, end_date, type, reminder, days, completed_dates) 
	VALUES (?,?,?,?,?,?,?,?,?)`

	if _, err := db.Exec(query,
		data.Name,
		data.UserId,
		data.Description,
		data.StartDate,
		data.EndDate,
		data.Type,
		data.Reminder,
		data.Days,
		data.CompletedDates,
	); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
