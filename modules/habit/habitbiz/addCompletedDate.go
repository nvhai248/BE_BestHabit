package habitbiz

import (
	"bestHabit/common"
	"bestHabit/modules/habit/habitmodel"
	"bestHabit/pubsub"
	"context"
)

type AddCompletedDateStore interface {
	UpdateCompletedDate(ctx context.Context, data *habitmodel.HabitUpdate, id int) error
	FindHabitById(ctx context.Context, id int) (*habitmodel.HabitFind, error)
}

type addCompletedDateBiz struct {
	store  AddCompletedDateStore
	pubsub pubsub.Pubsub
}

func NewAddCompletedDateBiz(store AddCompletedDateStore, pubsub pubsub.Pubsub) *addCompletedDateBiz {
	return &addCompletedDateBiz{store: store, pubsub: pubsub}
}

func (b *addCompletedDateBiz) AddCompletedDate(ctx context.Context, date *common.CompleteDate, id int) error {
	newData, err := b.store.FindHabitById(ctx, id)

	if err != nil {
		if err == common.ErrorNoRows {
			return common.ErrEntityNotFound(habitmodel.EntityName, err)
		}

		return err
	}

	newData.CompletedDates.AddDate(*date)

	if err := b.store.UpdateCompletedDate(ctx, &habitmodel.HabitUpdate{
		CompletedDates: newData.CompletedDates,
	}, id); err != nil {
		return err
	}
	return nil
}
