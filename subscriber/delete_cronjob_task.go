package subscriber

import (
	"bestHabit/component"
	"bestHabit/modules/cronnoticetask/cronnoticetaskstorage"
	"bestHabit/pubsub"
	"context"
)

func RunDeleteCronJobTaskAfterUserDeleteTask(appCtx component.AppContext) consumerJob {
	store := cronnoticetaskstorage.NewSQLStore(appCtx.GetMainDBConnection())

	return consumerJob{
		Title: "Create new cron job after user create new task!",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			userData := message.Data().(HasTaskInfoWhenDelete)

			cronJobs, err := store.GetListCronNoticeTask(ctx, userData.GetUserId(), userData.GetTaskId())

			if err != nil {
				return nil
			}

			for _, cronJob := range cronJobs {
				appCtx.GetCronJob().RemoveJob(cronJob.EntryId)
			}

			return store.DeleteCronNoticeTasks(ctx, userData.GetUserId(), userData.GetTaskId())
		},
	}
}
