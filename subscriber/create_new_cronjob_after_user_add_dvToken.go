package subscriber

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/cronnoticehabit/cronnoticehabitmodel"
	"bestHabit/modules/cronnoticehabit/cronnoticehabitstorage"
	"bestHabit/modules/cronnoticetask/cronnoticetaskmodel"
	"bestHabit/modules/cronnoticetask/cronnoticetaskstorage"
	"bestHabit/modules/habit/habitstorage"
	"bestHabit/modules/task/taskstorage"
	"bestHabit/pubsub"
	"context"
	"fmt"

	"github.com/robfig/cron/v3"
)

func RunCreateNewCronJobAfterUserAddDeviceToken(appCtx component.AppContext) consumerJob {
	store := cronnoticehabitstorage.NewSQLStore(appCtx.GetMainDBConnection())
	_store := cronnoticetaskstorage.NewSQLStore(appCtx.GetMainDBConnection())
	habitStore := habitstorage.NewSQLStore(appCtx.GetMainDBConnection())
	taskStore := taskstorage.NewSQLStore(appCtx.GetMainDBConnection())

	return consumerJob{
		Title: "Create new cron job after user add device token!",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			userData := message.Data().(HasUserIdAndUserLatestDvToken)

			if userData.GetLatestDvToken() == nil {
				return nil
			}

			sendNoticeFunc := func(deviceToken string, title, body string) error {
				return appCtx.GetSendNotification().SendNotification(deviceToken, title, body)
			}

			go func() {
				listHabit, err := habitStore.ListHabitByConditions(ctx, nil, nil, map[string]interface{}{"user_id": userData.GetUserId()})

				if err != nil {
					return
				}

				for _, h := range listHabit {
					entryIds, err := appCtx.GetCronJob().CreateNewJobs(*common.NewNotificationBasedHabit(userData.GetUserId(),
						h.Description,
						h.Name,
						h.StartDate,
						h.EndDate,
						h.Reminder,
						*h.Days), sendNoticeFunc,
						userData.GetLatestDvToken().DeviceToken,
						fmt.Sprintf("Time to do %s", h.Name),
						h.Description)

					if err != nil {
						return
					}

					for _, entryId := range entryIds {
						store.CreateNewCronNoticeHabit(ctx, &cronnoticehabitmodel.CronNoticeHabit{
							UserId:  userData.GetUserId(),
							HabitId: h.Id,
							EntryId: cron.EntryID(entryId),
						})
					}
				}
			}()

			go func() {
				listTask, err := taskStore.ListTaskByConditions(ctx, nil, nil, map[string]interface{}{"user_id": userData.GetUserId()})
				if err != nil {
					return
				}

				for _, t := range listTask {
					entryIds, err := appCtx.GetCronJob().CreateNewJobs(*common.NewNotificationBasedOnTask(
						userData.GetUserId(),
						t.Description,
						t.Name,
						t.Reminder),
						sendNoticeFunc,
						userData.GetLatestDvToken().DeviceToken,
						fmt.Sprintf("Time to do %s", t.Name),
						t.Name,
					)

					if err != nil {
						return
					}

					for _, entryId := range entryIds {
						_store.CreateNewCronNoticeTask(ctx, &cronnoticetaskmodel.CronNoticeTask{
							UserId:  userData.GetUserId(),
							TaskId:  t.Id,
							EntryId: cron.EntryID(entryId),
						})
					}
				}
			}()

			return nil
		},
	}
}
