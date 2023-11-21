package habitbiz

import (
	"bestHabit/common"
	"bestHabit/modules/habit/habitmodel"
	"context"
)

type FindHabitStorage interface {
	FindHabitById(ctx context.Context, id int) (*habitmodel.HabitFind, error)
}

type findHabitBiz struct {
	store FindHabitStorage
}

func NewFindHabitBiz(store FindHabitStorage) *findHabitBiz {
	return &findHabitBiz{store: store}
}

func (b *findHabitBiz) FindHabit(ctx context.Context, id int) (*habitmodel.HabitFind, error) {
	result, err := b.store.FindHabitById(ctx, id)

	if err != nil {
		if err == common.ErrorNoRows {
			return nil, common.ErrCannotGetEntity(habitmodel.EntityName, err)
		}
		return nil, err
	}

	if result.Status == 0 {
		return nil, common.ErrEntityDeleted(habitmodel.EntityName, err)
	}

	return result, nil
}
