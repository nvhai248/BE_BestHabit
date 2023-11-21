package habitbiz

import (
	"bestHabit/common"
	"bestHabit/modules/habit/habitmodel"
	"bestHabit/pubsub"
	"context"
)

type CreateHabitStore interface {
	Create(ctx context.Context, data *habitmodel.HabitCreate) error
}

type createHabitBiz struct {
	store  CreateHabitStore
	pubsub pubsub.Pubsub
}

func NewCreateHabitBiz(store CreateHabitStore, pubsub pubsub.Pubsub) *createHabitBiz {
	return &createHabitBiz{store: store, pubsub: pubsub}
}

func (b *createHabitBiz) CreateHabit(ctx context.Context, data *habitmodel.HabitCreate, userId int) error {
	if err := data.Validate(); err != nil {
		return err
	}

	data.UserId = userId

	if err := b.store.Create(ctx, data); err != nil {
		return err
	}

	b.pubsub.Publish(ctx, common.TopicUserCreateNewHabit, pubsub.NewMessage(data))
	return nil
}
