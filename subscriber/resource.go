package subscriber

import "bestHabit/common"

type HasUserId interface {
	GetUserId() int
}

type HasTaskInfoCreate interface {
	GetUserId() int
	GetDescription() string
	GetName() string
	GetReminderTime() string
}

type HasHabitInfoCreate interface {
	GetUserId() int
	GetDescription() string
	GetName() string
	GetReminderTime() string
	GetStartDate() string
	GetEndDate() string
	GetDays() *common.Days
}

type HasChallengeId interface {
	GetChallengeId() int
}

type HasTaskInfoWhenDelete interface {
	GetUserId() int
	GetTaskId() int
}

type HasHabitInfoWhenDelete interface {
	GetUserId() int
	GetHabitId() int
}
