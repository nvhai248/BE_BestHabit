package cronjob

import (
	"bestHabit/common"
	"fmt"

	"github.com/robfig/cron/v3"
)

type CronJobProvider interface {
	CreateNewJobs(notification common.Notification) ([]cron.EntryID, error)
	RemoveJob(entryId cron.EntryID) error
	StartJobs()
}

type CronJob struct {
	cronJob *cron.Cron
}

func NewCronJob() *CronJob {
	return &CronJob{cronJob: cron.New()}
}

// Function to check if the given day is a notification day according to Days in Notification
func (c *CronJob) isNotificationDay(weekday string, days common.Days) bool {

	for _, day := range days {
		if day.Weekday == weekday {
			return true
		}
	}
	return false
}

func (c *CronJob) CreateNewJobs(notification common.Notification) ([]cron.EntryID, error) {
	var entryIDs []cron.EntryID
	var err error

	if *notification.IsTask {
		t, err := common.ParseStringToTimestamp(*notification.ReminderTime)

		if err != nil {
			fmt.Println(err)
			return nil, common.ErrInternal(err)
		}

		_, month, day := t.Date()
		hour, minute, _ := t.Clock()

		entryID, _ := c.cronJob.AddFunc(fmt.Sprintf("%d %d %d %d *", minute, hour, day, month), func() {
			fmt.Println("Sending notification for task to user:", *notification.UserId)
			fmt.Println("Info:", notification)
		})

		entryIDs = append(entryIDs, entryID)
	} else {
		// Logic for habit notifications (if IsTask is false)
		// Assuming *notification.Days is not nil and correctly filled when IsTask is false

		// Parsing StartDate and EndDate
		startDate, err := common.ParseStringToDate(*notification.StartDate)
		if err != nil {
			fmt.Println(err)
			return nil, common.ErrInternal(err)
		}
		endDate, err := common.ParseStringToDate(*notification.EndDate)
		if err != nil {
			fmt.Println(err)
			return nil, common.ErrInternal(err)
		}

		// Splitting ReminderTime to get hours and minutes
		reminderTime, err := common.ParseStringToTime(*notification.ReminderTime)
		if err != nil {
			fmt.Println(err)
			return nil, common.ErrInternal(err)
		}
		hour, minute, _ := reminderTime.Clock()

		for date := *startDate; date.Before(*endDate) || date.Equal(*endDate); date = date.AddDate(0, 0, 1) {
			if c.isNotificationDay(date.Weekday().String(), *notification.Days) {
				entryId, _ := c.cronJob.AddFunc(fmt.Sprintf("%d %d %d %d *", minute, hour, date.Day(), date.Month()), func() {
					fmt.Println("Sending updated notification for habit to user:", *notification.UserId)
					fmt.Println("Info:", notification)
				})

				entryIDs = append(entryIDs, entryId)
			}
		}
	}

	if err != nil {
		return nil, err
	}

	return entryIDs, nil
}

func (c *CronJob) RemoveJob(entryId cron.EntryID) error {
	c.cronJob.Remove(cron.EntryID(entryId))
	return nil
}

func (c *CronJob) StartJobs() {
	c.cronJob.Start()
}
