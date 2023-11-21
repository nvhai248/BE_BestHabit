package taskstorage

import (
	"bestHabit/common"
	"bestHabit/modules/task/taskmodel"
	"context"
)

func (s *sqlStore) UpdateTaskInfo(ctx context.Context, newInfo *taskmodel.TaskUpdate, id int) error {
	db := s.db

	if _, err := db.Exec("UPDATE tasks SET name = ?, description = ?, deadline = ?, reminder = ? WHERE id = ?",
		newInfo.Name, newInfo.Description, newInfo.Deadline, newInfo.Reminder, id); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
