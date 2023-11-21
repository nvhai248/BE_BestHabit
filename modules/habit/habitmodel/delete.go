package habitmodel

type HabitDelete struct {
	Status int `json:"status" db:"status"`
}

func (HabitDelete) TableName() string {
	return Habit{}.TableName()
}
