package usermodel

import "bestHabit/common"

type UserUpdate struct {
	Phone    *string          `json:"phone" db:"phone"`
	Name     *string          `json:"name" db:"name"`
	Avatar   *common.Image    `json:"avatar" db:"avatar"`
	Settings *common.Settings `json:"settings" db:"settings"`
}

func (UserUpdate) TableName() string {
	return UserCreate{}.TableName()
}
