package taskstorage

import (
	"bestHabit/common"
	"bestHabit/modules/task/taskmodel"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *taskmodel.TaskCreate) error {
	db := s.db

	query := `INSERT INTO tasks (name, user_id, description, deadline, reminder, status) 
	VALUES (?,?,?,?,?,?)`

	deadlineDate, err := common.ParseStringToDate(data.Deadline)
	if err != nil {
		return common.ErrInternal(err)
	}

	reminderTime, err := common.ParseStringToTimestamp(data.Reminder)
	if err != nil {
		return common.ErrInternal(err)
	}

	if _, err := db.Exec(query, data.Name, data.UserId, data.Description, deadlineDate, reminderTime, data.Status); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
