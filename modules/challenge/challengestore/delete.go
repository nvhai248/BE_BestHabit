package challengestore

import (
	"bestHabit/common"
	"context"
)

func (s *sqlStore) SoftDelete(ctx context.Context, id int) error {
	db := s.db

	if _, err := db.Exec("UPDATE challenges SET status = 0 WHERE id = ?", id); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
