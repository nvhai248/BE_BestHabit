package subscriber

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/cronnoticehabit/cronnoticehabitmodel"
	"bestHabit/modules/cronnoticehabit/cronnoticehabitstorage"
	"bestHabit/modules/habit/habitstorage"
	"bestHabit/modules/user/userstorage"
	"bestHabit/pubsub"
	"context"
	"fmt"

	"github.com/robfig/cron/v3"
)

func RunCreateNewCronJobHabitAfterUserAddNewHabit(appCtx component.AppContext) consumerJob {
	store := cronnoticehabitstorage.NewSQLStore(appCtx.GetMainDBConnection())
	habitStore := habitstorage.NewSQLStore(appCtx.GetMainDBConnection())
	userStore := userstorage.NewSQLStore(appCtx.GetMainDBConnection())

	return consumerJob{
		Title: "Create new cron job after user create new task!",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			habitData := message.Data().(HasHabitInfoCreate)

			if habitData.GetReminderTime() == "" {
				return nil
			}

			userData, err := userStore.FindById(ctx, habitData.GetUserId())
			if err != nil {
				return nil
			}

			sendNoticeFunc := func(deviceToken string, title, body string) error {
				return appCtx.GetSendNotification().SendNotification(deviceToken, title, body)
			}

			for _, token := range *userData.DeviceTokens {

				entryIds, err := appCtx.GetCronJob().CreateNewJobs(*common.NewNotificationBasedHabit(habitData.GetUserId(),
					habitData.GetDescription(),
					habitData.GetName(),
					habitData.GetStartDate(),
					habitData.GetEndDate(),
					habitData.GetReminderTime(),
					*habitData.GetDays()), sendNoticeFunc,
					token.DeviceToken,
					fmt.Sprintf("Time to do %s", habitData.GetName()),
					habitData.GetDescription())

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
			}

			return nil
		},
	}
}
