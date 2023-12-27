package subscriber

import (
	"bestHabit/component"
	"bestHabit/modules/cronnoticehabit/cronnoticehabitstorage"
	"bestHabit/pubsub"
	"context"
)

func RunDeleteCronJobHabitAfterUserDeleteHabit(appCtx component.AppContext) consumerJob {
	store := cronnoticehabitstorage.NewSQLStore(appCtx.GetMainDBConnection())

	return consumerJob{
		Title: "Delete cron job after user delete Habit!",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			userData := message.Data().(HasHabitInfoWhenDelete)

			cronJobs, err := store.GetListCronNoticeHabit(ctx, userData.GetUserId(), userData.GetHabitId())

			if err != nil {
				return nil
			}

			for _, cronJob := range cronJobs {
				appCtx.GetCronJob().RemoveJob(cronJob.EntryId)
			}

			return store.DeleteCronNoticeHabits(ctx, userData.GetUserId(), userData.GetHabitId())
		},
	}
}
