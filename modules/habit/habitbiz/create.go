package habitbiz

import (
	"bestHabit/common"
	"bestHabit/component/cronjob"
	"bestHabit/modules/habit/habitmodel"
	"bestHabit/pubsub"
	"context"
	"fmt"
)

type CreateHabitStore interface {
	Create(ctx context.Context, data *habitmodel.HabitCreate) error
}

type createHabitBiz struct {
	store   CreateHabitStore
	pubsub  pubsub.Pubsub
	cronJob cronjob.CronJobProvider
}

func NewCreateHabitBiz(store CreateHabitStore, pubsub pubsub.Pubsub, cronJob cronjob.CronJobProvider) *createHabitBiz {
	return &createHabitBiz{store: store, pubsub: pubsub, cronJob: cronJob}
}

func (b *createHabitBiz) CreateHabit(ctx context.Context, data *habitmodel.HabitCreate, userId int) error {
	if err := data.Validate(); err != nil {
		return err
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

	// create a new cron job
	go func() {
		defer common.AppRecover()
		entryIds, _ := b.cronJob.CreateNewJobs(*common.NewNotificationBasedHabit(userId,
			data.Description,
			data.Name,
			data.StartDate,
			data.EndDate,
			data.Reminder,
			*data.Days))
		fmt.Println(entryIds)
	}()
	return nil
}
