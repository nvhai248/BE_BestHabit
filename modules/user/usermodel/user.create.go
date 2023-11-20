package usermodel

import (
	"bestHabit/common"
	"errors"
	"strings"
)

const EntityName = "User"

type UserCreate struct {
	Email    string           `json:"email" db:"email"`
	Phone    string           `json:"phone" db:"phone"`
	Password string           `json:"-" db:"password"`
	Name     string           `json:"name" db:"name"`
	FbID     string           `json:"-" db:"fb_id"`
	GgID     string           `json:"-" db:"gg_id"`
	Salt     string           `json:"-" db:"salt"`
	Avatar   *common.Image    `json:"avatar" db:"avatar"`
	Settings *common.Settings `json:"settings" db:"settings"`
	Role     string           `json:"role" db:"role"`
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
