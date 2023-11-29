package taskstorage

import (
	"bestHabit/common"
	"context"
)

func (s *sqlStore) DeleteTask(ctx context.Context, id int) error {
	db := s.db

	if _, err := db.Exec("UPDATE tasks SET status = 'deleted' WHERE id = ?", id); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
