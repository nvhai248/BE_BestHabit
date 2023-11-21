package habitbiz

import (
	"bestHabit/common"
	"bestHabit/modules/habit/habitmodel"
	"context"
)

type SoftDeleteHabitStorage interface {
	SoftDeleteHabit(ctx context.Context, id int) error
	FindHabitById(ctx context.Context, id int) (*habitmodel.HabitFind, error)
}

type softDeleteHabitBiz struct {
	store SoftDeleteHabitStorage
}

func NewSoftDeleteHabitBiz(store SoftDeleteHabitStorage) *softDeleteHabitBiz {
	return &softDeleteHabitBiz{store: store}
}

func (b *softDeleteHabitBiz) SoftDeleteHabit(ctx context.Context, id int) error {
	result, err := b.store.FindHabitById(ctx, id)

	if err != nil {
		if err == common.ErrorNoRows {
			return common.ErrEntityNotFound(habitmodel.EntityName, err)
		}

		return err
	}

	if result.Status == 0 {
		return common.ErrEntityDeleted(habitmodel.EntityName, err)
	}

	if err := b.store.SoftDeleteHabit(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(habitmodel.EntityName, err)
	}

	return nil
}
