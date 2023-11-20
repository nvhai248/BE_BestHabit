package taskmodel

type TaskCreate struct {
	UserId      int    `json:"-" db:"user_id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Deadline    string `json:"deadline" db:"deadline"`
	Reminder    string `json:"reminder" db:"reminder"`
}

func (TaskCreate) TableName() string {
	return Task{}.TableName()
}

func (t *TaskCreate) GetUserId() int {
	return t.UserId
}

func (t *TaskCreate) Validate() error {
	if t.Name == "" {
		return ErrNameNotBeBlank
	}

	if t.Deadline == "" {
		return ErrDeadlineNotBeBlank
	}

	return nil
}
