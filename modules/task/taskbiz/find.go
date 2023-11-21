package taskbiz

import (
	"bestHabit/common"
	"bestHabit/modules/task/taskmodel"
	"context"
)

type FindTaskStorage interface {
	FindTaskById(ctx context.Context, id int) (*taskmodel.TaskFind, error)
}

type findTaskBiz struct {
	store FindTaskStorage
}

func NewFindTaskBiz(store FindTaskStorage) *findTaskBiz {
	return &findTaskBiz{store: store}
}

func (b *findTaskBiz) FindTask(ctx context.Context, id int) (*taskmodel.TaskFind, error) {
	result, err := b.store.FindTaskById(ctx, id)

	if err != nil {
		if err == common.ErrorNoRows {
			return nil, common.ErrCannotGetEntity(taskmodel.EntityName, err)
		}
		return nil, err
	}

	if result.Status == "deleted" {
		return nil, common.ErrEntityDeleted(taskmodel.EntityName, err)
	}

	return result, nil
}
