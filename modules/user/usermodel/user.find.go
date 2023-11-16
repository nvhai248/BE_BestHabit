package usermodel

import (
	"bestHabit/common"
)

type UserFind struct {
	common.SQLModel `json: ", inline"`

	Email          *string          `json:"email" db:"email"`
	Phone          *string          `json:"phone" db:"phone"`
	Password       *string          `json:"-" db:"password"`
	Name           *string          `json:"name" db:"name"`
	FbID           *string          `json:"-" db:"fb_id"`
	GgID           *string          `json:"-" db:"gg_id"`
	Salt           *string          `json:"-" db:"salt"`
	Avatar         *common.Image    `json:"avatar" db:"avatar"`
	Level          int              `json:"level" db:"level"`
	Experience     int              `json:"experience" db:"experience"`
	Settings       *common.Settings `json:"settings" db:"settings"`
	Role           *string          `json:"role" db:"role"`
	HabitCount     int              `json:"habit_count" db:"habit_count"`
	TaskCount      int              `json:"task_count" db:"task_count"`
	ChallengeCount int              `json:"challenge_count" db:"challenge_count"`
}

func (UserFind) TableName() string {
	return UserCreate{}.TableName()
}
