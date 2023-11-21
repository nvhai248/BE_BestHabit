package habitstorage

import (
	"bestHabit/common"
	"bestHabit/modules/habit/habitmodel"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *habitmodel.HabitCreate) error {
	db := s.db

	query := `INSERT INTO habits (name, user_id, description, start_date, end_date, type, reminder, days) 
	VALUES (?,?,?,?,?)`

	startDate, err := common.ParseStringToDate(data.StartDate)
	if err != nil {
		return common.ErrInternal(err)
	}

	endDate, err := common.ParseStringToDate(data.EndDate)
	if err != nil {
		return common.ErrInternal(err)
	}

	reminderTime, err := common.ParseStringToTime(data.Reminder)
	if err != nil {
		return common.ErrInternal(err)
	}

	if _, err := db.Exec(query,
		data.Name,
		data.UserId,
		data.Description,
		startDate,
		endDate,
		data.Type,
		reminderTime,
		data.Days,
	); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
