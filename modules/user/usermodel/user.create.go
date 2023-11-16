package usermodel

import (
	"bestHabit/common"
	"errors"
	"strings"
)

const EntityName = "User"

type UserCreate struct {
	common.SQLModel `json:", inline"`

	Email    string           `json:"email" db:"email"`
	Phone    string           `json:"phone" db:"phone"`
	Password string           `json:"password" db:"password"`
	Name     string           `json:"name" db:"name"`
	FbID     string           `json:"fb_id" db:"fb_id"`
	GgID     string           `json:"gg_id" db:"gg_id"`
	Salt     string           `json:"salt" db:"salt"`
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

func (user *UserCreate) Mask(isAdminOrOwner bool) {
	user.GenUID(common.DbTypeUser)
}
