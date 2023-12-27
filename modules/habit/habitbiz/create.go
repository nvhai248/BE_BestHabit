package habitbiz

import (
	"bestHabit/common"
	"bestHabit/modules/habit/habitmodel"
	"bestHabit/pubsub"
	"context"
)

type CreateHabitStore interface {
	Create(ctx context.Context, data *habitmodel.HabitCreate) error
	FindHabitByInformation(ctx context.Context, userId int, name string) (*habitmodel.HabitFind, error)
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

	if _, err := b.store.FindHabitByInformation(ctx, userId, data.Name); err == nil {
		return habitmodel.ErrNameAlreadyUsed
	}

	data.UserId = userId
	data.CompletedDates = &common.CompleteDates{}
	data.CompletedDates.Init()

	if data.Days == nil {
		data.Days = &common.Days{}
		data.Days.Init()
	}

	if data.Target == nil {
		data.Target = common.NewDefaultTarget()
	}

	if err := b.store.Create(ctx, data); err != nil {
		return err
	}
	go func() {
		defer common.AppRecover()
		b.pubsub.Publish(ctx, common.TopicUserCreateNewHabit, pubsub.NewMessage(data))
	}()
	return nil
}
