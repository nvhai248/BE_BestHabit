package taskbiz

import (
	"bestHabit/common"
	"bestHabit/modules/task/taskmodel"
	"context"
)

type UpdateTaskStorage interface {
	FindTaskById(ctx context.Context, id int) (*taskmodel.TaskFind, error)
	UpdateTaskInfo(ctx context.Context, newInfo *taskmodel.TaskUpdate, id int) error
}

type updateTaskBiz struct {
	store UpdateTaskStorage
}

func NewUpdateTaskBiz(store UpdateTaskStorage) *updateTaskBiz {
	return &updateTaskBiz{store: store}
}

func (b *updateTaskBiz) Update(ctx context.Context, newInfo *taskmodel.TaskUpdate, id int) error {
	oldData, err := b.store.FindTaskById(ctx, id)

	if err != nil {
		if err == common.ErrorNoRows {
			return common.ErrEntityNotFound(taskmodel.EntityName, err)
		}

		return err
	}

	if oldData.Status == "deleted" {
		return common.ErrEntityDeleted(taskmodel.EntityName, err)
	}

	if newInfo.Status == nil {
		newInfo.Status = &oldData.Status
	}

	if newInfo.Name == nil {
		newInfo.Name = &oldData.Name
	}

	if newInfo.Description == nil {
		newInfo.Description = &oldData.Description
	}

	if newInfo.Deadline == nil {
		newInfo.Deadline = &oldData.Deadline
	}

	if newInfo.Reminder == nil {
		newInfo.Reminder = &oldData.Reminder
	}

	err = b.store.UpdateTaskInfo(ctx, newInfo, id)

	if err != nil {
		return common.ErrCannotUpdateEntity(taskmodel.EntityName, err)
	}

	return nil
}
