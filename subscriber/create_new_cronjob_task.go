package subscriber

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/cronnoticetask/cronnoticetaskmodel"
	"bestHabit/modules/cronnoticetask/cronnoticetaskstorage"
	"bestHabit/modules/task/taskstorage"
	"bestHabit/pubsub"
	"context"

	"github.com/robfig/cron/v3"
)

func RunCreateNewCronJobTaskAfterUserAddNewTask(appCtx component.AppContext) consumerJob {
	store := cronnoticetaskstorage.NewSQLStore(appCtx.GetMainDBConnection())
	taskStore := taskstorage.NewSQLStore(appCtx.GetMainDBConnection())

	return consumerJob{
		Title: "Create new cron job after user create new task!",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			userData := message.Data().(HasTaskInfoCreate)
			entryIds, err := appCtx.GetCronJob().CreateNewJobs(*common.NewNotificationBasedOnTask(userData.GetUserId(),
				userData.GetDescription(),
				userData.GetName(),
				userData.GetReminderTime()))

			if err != nil {
				return err
			}

			task, err := taskStore.FindTaskByInformation(ctx, userData.GetUserId(), userData.GetName())

			if err != nil {
				return nil
			}

			for _, entryId := range entryIds {
				store.CreateNewCronNoticeTask(ctx, &cronnoticetaskmodel.CronNoticeTask{
					UserId:  userData.GetUserId(),
					TaskId:  task.Id,
					EntryId: cron.EntryID(entryId),
				})
			}

			return nil
		},
	}
}
