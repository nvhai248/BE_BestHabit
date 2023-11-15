package usermodel

import (
	"bestHabit/common"
	"database/sql"
	"errors"
	"strings"
)

const EntityName = "User"

type UserCreate struct {
	Email    string          `json:"email" db:"email"`
	Phone    string          `json:"phone" db:"phone"`
	Password string          `json:"password" db:"password"`
	Name     string          `json:"name" db:"name"`
	FbID     sql.NullString  `json:"fb_id" db:"fb_id"`
	GgID     sql.NullString  `json:"gg_id" db:"gg_id"`
	Salt     sql.NullString  `json:"salt" db:"salt"`
	Avatar   common.Image    `json:"avatar" db:"avatar"`
	Settings common.Settings `json:"settings" db:"settings"`
	Role     string          `json:"role" db:"role"`
}

func (UserCreate) TableName() string {
	return "users"
}

func (u *UserCreate) validate() error {
	u.Name = strings.TrimSpace(u.Name)

	if len(u.Name) == 0 {
		return errors.New("student name cannot be blank!")
	}

	return nil
}
