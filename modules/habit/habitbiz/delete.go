package habitbiz

import (
	"bestHabit/common"
	"bestHabit/modules/habit/habitmodel"
	"bestHabit/pubsub"
	"context"
)

type SoftDeleteHabitStorage interface {
	SoftDeleteHabit(ctx context.Context, id int) error
	FindHabitById(ctx context.Context, id int) (*habitmodel.HabitFind, error)
}

type softDeleteHabitBiz struct {
	store  SoftDeleteHabitStorage
	pubsub pubsub.Pubsub
}

func NewSoftDeleteHabitBiz(store SoftDeleteHabitStorage, pubsub pubsub.Pubsub) *softDeleteHabitBiz {
	return &softDeleteHabitBiz{store: store, pubsub: pubsub}
}

func (b *softDeleteHabitBiz) SoftDeleteHabit(ctx context.Context, id int) error {
	habit, err := b.store.FindHabitById(ctx, id)

	if err != nil {
		if err == common.ErrorNoRows {
			return common.ErrEntityNotFound(habitmodel.EntityName, err)
		}

		return err
	}

	if habit.Status == 0 {
		return common.ErrEntityDeleted(habitmodel.EntityName, err)
	}

	if err := b.store.SoftDeleteHabit(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(habitmodel.EntityName, err)
	}

	go func() {
		defer common.AppRecover()
		b.pubsub.Publish(ctx, common.TopicUserDeleteHabit, pubsub.NewMessage(habit))
	}()

	return nil
}
