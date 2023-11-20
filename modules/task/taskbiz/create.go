package taskbiz

import (
	"bestHabit/modules/task/taskmodel"
	"context"
)

type CreateTaskStore interface {
	Create(ctx context.Context, data *taskmodel.TaskCreate) error
}

type createTaskBiz struct {
	store CreateTaskStore
}

func NewCreateTaskBiz(store CreateTaskStore) *createTaskBiz {
	return &createTaskBiz{store: store}
}

func (b *createTaskBiz) CreateTask(ctx context.Context, data *taskmodel.TaskCreate, userId int) error {
	if err := data.Validate(); err != nil {
		return err
	}

	data.Status = "pending"
	data.UserId = userId

	if err := b.store.Create(ctx, data); err != nil {
		return err
	}

	return nil
}
