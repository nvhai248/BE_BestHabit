package taskbiz

import (
	"bestHabit/common"
	"bestHabit/modules/task/taskmodel"
	"context"
)

type ListTaskStorage interface {
	ListTaskByConditions(ctx context.Context,
		filter *taskmodel.TaskFilter,
		paging *common.Paging,
		conditions map[string]interface{}) ([]taskmodel.Task, error)
}

type listTaskBiz struct {
	store ListTaskStorage
}

func NewListTaskBiz(store ListTaskStorage) *listTaskBiz {
	return &listTaskBiz{store: store}
}

func (b *listTaskBiz) ListTask(ctx context.Context,
	filter *taskmodel.TaskFilter,
	paging *common.Paging,
	conditions map[string]interface{}) ([]taskmodel.Task, error) {

	task, err := b.store.ListTaskByConditions(ctx, filter, paging, conditions)

	if err != nil {
		return nil, err
	}

	return task, nil
}
