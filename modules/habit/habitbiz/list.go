package habitbiz

import (
	"bestHabit/common"
	"bestHabit/modules/habit/habitmodel"
	"context"
)

type ListHabitStorage interface {
	ListHabitByConditions(ctx context.Context,
		filter *habitmodel.HabitFilter,
		paging *common.Paging,
		conditions map[string]interface{}) ([]habitmodel.Habit, error)
}

type listTaskBiz struct {
	store ListHabitStorage
}

func NewListHabitBiz(store ListHabitStorage) *listTaskBiz {
	return &listTaskBiz{store: store}
}

func (b *listTaskBiz) ListHabit(ctx context.Context,
	filter *habitmodel.HabitFilter,
	paging *common.Paging,
	conditions map[string]interface{}) ([]habitmodel.Habit, error) {

	habits, err := b.store.ListHabitByConditions(ctx, filter, paging, conditions)

	if err != nil {
		return nil, err
	}

	return habits, nil
}
