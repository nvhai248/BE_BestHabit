package cronjob

import (
	"bestHabit/common"
	"fmt"

	"github.com/robfig/cron/v3"
)

// Add a job to Group
func AddJobToGroup(group *cron.Cron, notification common.Notification) (cron.EntryID, error) {
	return group.AddFunc(*notification.ReminderTime, func() {
		fmt.Println("Adding job to group and sending notification to user:", notification.UserId)
	})
}

// Remove Job from group
func RemoveJobFromGroup(group *cron.Cron, entryID cron.EntryID) {
	group.Remove(entryID)
}

// Run concurrently jobs in group
func RunAllJobs(group *cron.Cron) {
	group.Start()
}
