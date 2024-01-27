package subscriber

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/cronnoticetask/cronnoticetaskmodel"
	"bestHabit/modules/cronnoticetask/cronnoticetaskstorage"
	"bestHabit/modules/task/taskstorage"
	"bestHabit/modules/user/userstorage"
	"bestHabit/pubsub"
	"context"
	"fmt"

	"github.com/robfig/cron/v3"
)

func RunCreateNewCronJobTaskAfterUserAddNewTask(appCtx component.AppContext) consumerJob {
	store := cronnoticetaskstorage.NewSQLStore(appCtx.GetMainDBConnection())
	taskStore := taskstorage.NewSQLStore(appCtx.GetMainDBConnection())
	userStore := userstorage.NewSQLStore(appCtx.GetMainDBConnection())

	return consumerJob{
		Title: "Create new cron job after user create new task!",
		Hld: func(ctx context.Context, message *pubsub.Message) error {

			taskData := message.Data().(HasTaskInfoCreate)

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
					taskData.GetDescription(),
				)

				if err != nil {
					return err
				}

				task, err := taskStore.FindTaskByInformation(ctx, taskData.GetUserId(), taskData.GetName())

				if err != nil {
					return nil
				}

				for _, entryId := range entryIds {
					store.CreateNewCronNoticeTask(ctx, &cronnoticetaskmodel.CronNoticeTask{
						UserId:  taskData.GetUserId(),
						TaskId:  task.Id,
						EntryId: cron.EntryID(entryId),
					})
				}
			}

			return nil
		},
	}
}
