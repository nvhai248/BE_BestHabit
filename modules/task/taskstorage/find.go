package taskstorage

import (
	"bestHabit/common"
	"bestHabit/modules/task/taskmodel"
	"context"
	"database/sql"
)

func (s *sqlStore) FindTaskById(ctx context.Context, id int) (*taskmodel.TaskFind, error) {
	db := s.db

	var task taskmodel.TaskFind
	if err := db.Get(&task, "SELECT * FROM tasks WHERE id = ?", id); err != nil {
		if err == sql.ErrNoRows {
			return nil, common.ErrorNoRows
		}

		return nil, common.ErrDB(err)
	}

	return &task, nil
}

func (s *sqlStore) FindTaskByInformation(ctx context.Context, userId int, name string) (*taskmodel.TaskFind, error) {
	db := s.db

	var task taskmodel.TaskFind
	if err := db.Get(&task, "SELECT * FROM tasks WHERE user_id = ? AND name = ? AND status <> 'deleted'", userId, name); err != nil {
		if err == sql.ErrNoRows {
			return nil, common.ErrorNoRows
		}

		return nil, common.ErrDB(err)
	}

	return &task, nil
}
