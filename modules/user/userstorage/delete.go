package userstorage

import (
	"bestHabit/common"
	"context"
)

func (s *sqlStore) DeleteUser(ctx context.Context, userId int) error {
	db := s.db

	if _, err := db.Exec("UPDATE users SET status = 0 WHERE id = ?", userId); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
