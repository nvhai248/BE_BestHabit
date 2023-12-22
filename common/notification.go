package common

type Notification struct {
	UserId       *int    `json:"user_id"`
	Description  *string `json:"description"`
	Name         *string `json:"name"`
	ReminderTime *string `json:"reminder_time"` // time remind
	IsTask       *bool   `json:"is_task"`       // task => true
	StartDate    *string `json:"start_date"`    // start date (only habit)
	EndDate      *string `json:"end_date"`      // end date (only habit)
	Days         *Days   `json:"days"`          // weekdays (only habit)
}

func NewNotificationBasedOnTask(userId int, description string, name string, reminderTime string) *Notification {
	isTask := true
	return &Notification{
		UserId:       &userId,
		Description:  &description,
		Name:         &name,
		ReminderTime: &reminderTime,
		StartDate:    nil,
		EndDate:      nil,
		Days:         nil,
		IsTask:       &isTask,
	}
}

func NewNotificationBasedHabit(userId int,
	description string,
	name string,
	startDate string,
	endDate string,
	reminderTime string,
	days Days) *Notification {
	isTask := false
	return &Notification{
		UserId:       &userId,
		Description:  &description,
		Name:         &name,
		ReminderTime: &reminderTime,
		StartDate:    &startDate,
		EndDate:      &endDate,
		Days:         &days,
		IsTask:       &isTask,
	}
}
