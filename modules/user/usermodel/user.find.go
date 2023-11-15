package usermodel

import (
	"bestHabit/common"
	"database/sql"
)

type UserFind struct {
	common.SQLModel `json: ", inline"`

	Email          string         `json:"email" db:"email"`
	Phone          string         `json:"phone" db:"phone"`
	Password       string         `json:"password" db:"-"`
	Name           string         `json:"name" db:"name"`
	FbID           sql.NullString `json:"fb_id" db:"-"`
	GgID           sql.NullString `json:"gg_id" db:"-"`
	Salt           sql.NullString `json:"salt" db:"-"`
	Avatar         sql.NullString `json:"avatar" db:"avatar"`
	Level          int            `json:"level" db:"level"`
	Experience     int            `json:"experience" db:"experience"`
	Settings       sql.NullString `json:"settings" db:"settings"`
	Role           string         `json:"role" db:"role"`
	HabitCount     int            `json:"habit_count" db:"habit_count"`
	TaskCount      int            `json:"task_count" db:"task_count"`
	ChallengeCount int            `json:"challenge_count" db:"challenge_count"`
}

func (UserFind) TableName() string {
	return UserCreate{}.TableName()
}
