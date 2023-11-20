package userstorage

import (
	"bestHabit/common"
	"context"
)

func (s *sqlStore) IncreaseHabitCount(ctx context.Context, userId int) error {
	db := s.db

	if _, err := db.Exec("UPDATE users SET habit_count = habit_count + 1 WHERE id = ?", userId); err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) IncreaseTaskCount(ctx context.Context, userId int) error {
	db := s.db

	if _, err := db.Exec("UPDATE users SET task_count = task_count + 1 WHERE id = ?", userId); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
