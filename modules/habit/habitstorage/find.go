package habitstorage

import (
	"bestHabit/common"
	"bestHabit/modules/habit/habitmodel"
	"context"
	"database/sql"
)

func (s *sqlStore) FindHabitById(ctx context.Context, id int) (*habitmodel.HabitFind, error) {
	db := s.db

	var habit habitmodel.HabitFind
	if err := db.Get(&habit, "SELECT * FROM habits WHERE id = ?", id); err != nil {
		if err == sql.ErrNoRows {
			return nil, common.ErrorNoRows
		}

		return nil, common.ErrDB(err)
	}

	return &habit, nil
}
