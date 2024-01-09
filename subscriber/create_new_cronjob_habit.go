package subscriber

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/cronnoticehabit/cronnoticehabitmodel"
	"bestHabit/modules/cronnoticehabit/cronnoticehabitstorage"
	"bestHabit/modules/habit/habitstorage"
	"bestHabit/pubsub"
	"context"

	"github.com/robfig/cron/v3"
)

func RunCreateNewCronJobHabitAfterUserAddNewHabit(appCtx component.AppContext) consumerJob {
	store := cronnoticehabitstorage.NewSQLStore(appCtx.GetMainDBConnection())
	habitStore := habitstorage.NewSQLStore(appCtx.GetMainDBConnection())

	return consumerJob{
		Title: "Create new cron job after user create new task!",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			habitData := message.Data().(HasHabitInfoCreate)

			if habitData.GetReminderTime() == "" {
				return nil
			}

			entryIds, err := appCtx.GetCronJob().CreateNewJobs(*common.NewNotificationBasedHabit(habitData.GetUserId(),
				habitData.GetDescription(),
				habitData.GetName(),
				habitData.GetStartDate(),
				habitData.GetEndDate(),
				habitData.GetReminderTime(),
				*habitData.GetDays()))

			if err != nil {
				return err
			}

			habit, err := habitStore.FindHabitByInformation(ctx, habitData.GetUserId(), habitData.GetName())

			if err != nil {
				return nil
			}

			for _, entryId := range entryIds {
				store.CreateNewCronNoticeHabit(ctx, &cronnoticehabitmodel.CronNoticeHabit{
					UserId:  habitData.GetUserId(),
					HabitId: habit.Id,
					EntryId: cron.EntryID(entryId),
				})
			}

			return nil
		},
	}
}
