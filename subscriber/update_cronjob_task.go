package subscriber

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/cronnoticetask/cronnoticetaskmodel"
	"bestHabit/modules/cronnoticetask/cronnoticetaskstorage"
	"bestHabit/modules/user/userstorage"
	"bestHabit/pubsub"
	"context"
	"fmt"

	"github.com/robfig/cron/v3"
)

func RunUpdateCronJobTaskAfterUserUpdateTask(appCtx component.AppContext) consumerJob {
	store := cronnoticetaskstorage.NewSQLStore(appCtx.GetMainDBConnection())
	userStore := userstorage.NewSQLStore(appCtx.GetMainDBConnection())

	return consumerJob{
		Title: "Update cron job after user update task!",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			// get Data from message
			taskData := message.Data().(HasTaskInfoUpdate)

			// remove cron jobs
			cronJobs, err := store.GetListCronNoticeTask(ctx, taskData.GetUserId(), taskData.GetTaskId())

			if err != nil {
				return nil
			}

			for _, cronJob := range cronJobs {
				appCtx.GetCronJob().RemoveJob(cronJob.EntryId)
			}

			// remove data in db
			store.DeleteCronNoticeTasks(ctx, taskData.GetUserId(), taskData.GetTaskId())

			userData, err := userStore.FindById(ctx, taskData.GetUserId())
			if err != nil {
				return nil
			}

			sendNoticeFunc := func(deviceToken string, title, body string) error {
				return appCtx.GetSendNotification().SendNotification(deviceToken, title, body)
			}

			for _, token := range *userData.DeviceTokens {
				entryIds, err := appCtx.GetCronJob().CreateNewJobs(*common.NewNotificationBasedOnTask(taskData.GetUserId(),
					taskData.GetDescription(),
					taskData.GetName(),
					taskData.GetReminderTime()),
					sendNoticeFunc,
					token.DeviceToken,
					fmt.Sprintf("Time to do %s", taskData.GetName()),
					taskData.GetDescription())

				if err != nil {
					fmt.Println(err)
					return err
				}

				for _, entryId := range entryIds {
					store.CreateNewCronNoticeTask(ctx, &cronnoticetaskmodel.CronNoticeTask{
						UserId:  taskData.GetUserId(),
						TaskId:  taskData.GetTaskId(),
						EntryId: cron.EntryID(entryId),
					})
				}
			}

			return nil
		},
	}
}
