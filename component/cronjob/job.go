package cronjob

import (
	"bestHabit/common"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
)

// Create new cron job
func CreateCronJob(notification common.Notification) (cron.EntryID, error) {
	c := cron.New()

	var entryID cron.EntryID
	var err error

	if *notification.IsTask {
		entryID, err = c.AddFunc(*notification.ReminderTime, func() {
			fmt.Println("Sending notification for task to user:", *notification.UserId)
			fmt.Println("Info:", notification)
		})
	} else {
		// Logic for habit notifications (if IsTask is false)
		// Assuming *notification.Days is not nil and correctly filled when IsTask is false

		// Parsing StartDate and EndDate
		startDate, _ := time.Parse("2006-01-02", *notification.StartDate)
		endDate, _ := time.Parse("2006-01-02", *notification.EndDate)

		// Splitting ReminderTime to get hours and minutes
		reminderTimeParts := strings.Split(*notification.ReminderTime, ":")
		hour, _ := strconv.Atoi(reminderTimeParts[0])
		minute, _ := strconv.Atoi(reminderTimeParts[1])

		for date := startDate; date.Before(endDate) || date.Equal(endDate); date = date.AddDate(0, 0, 1) {
			if isNotificationDay(date.Weekday().String(), *notification.Days) {
				entryID, err = c.AddFunc(fmt.Sprintf("%d %d %d %d *", minute, hour, date.Day(), date.Month()), func() {
					fmt.Println("Sending notification for habit to user:", *notification.UserId)
					fmt.Println("Info:", notification)
				})
			}
		}
	}

	if err != nil {
		return 0, err
	}

	// Start cron job
	c.Start()

	return entryID, nil
}

// Function to check if the given day is a notification day according to Days in Notification
func isNotificationDay(weekday string, days common.Days) bool {
	for _, day := range days {
		if day.Weekday == weekday {
			return true
		}
	}
	return false
}

// Update cron job
func UpdateCronJob(entryID cron.EntryID, newDateTime string, notification common.Notification) error {
	c := cron.New()

	c.Remove(entryID)

	var err error
	if *notification.IsTask {
		_, err = c.AddFunc(newDateTime, func() {
			fmt.Println("Sending updated notification for task to user:", *notification.UserId)
			fmt.Println("Info:", notification)
		})
	} else {
		// Logic for habit notifications (if IsTask is false)
		// Assuming *notification.Days is not nil and correctly filled when IsTask is false

		// Parsing StartDate and EndDate
		startDate, _ := time.Parse("2006-01-02", *notification.StartDate)
		endDate, _ := time.Parse("2006-01-02", *notification.EndDate)

		// Splitting ReminderTime to get hours and minutes
		reminderTimeParts := strings.Split(*notification.ReminderTime, ":")
		hour, _ := strconv.Atoi(reminderTimeParts[0])
		minute, _ := strconv.Atoi(reminderTimeParts[1])

		for date := startDate; date.Before(endDate) || date.Equal(endDate); date = date.AddDate(0, 0, 1) {
			if isNotificationDay(date.Weekday().String(), *notification.Days) {
				_, err = c.AddFunc(fmt.Sprintf("%d %d %d %d *", minute, hour, date.Day(), date.Month()), func() {
					fmt.Println("Sending updated notification for habit to user:", *notification.UserId)
					fmt.Println("Info:", notification)
				})
			}
		}
	}

	if err != nil {
		return err
	}

	c.Start()

	return nil
}

// Remove cron job
func RemoveCronJob(entryID cron.EntryID) {
	c := cron.New()

	c.Remove(entryID)
}
