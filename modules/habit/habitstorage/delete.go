package habitstorage

import (
	"bestHabit/common"
	"context"
	"database/sql"
)

func (s *sqlStore) SoftDeleteHabit(ctx context.Context, id int) error {
	db := s.db

	if _, err := db.Exec("UPDATE habits SET status = 0 where id = ?", id); err != nil {
		if err == sql.ErrNoRows {
			return common.ErrorNoRows
		}

		return common.ErrDB(err)
	}

	return nil
}
