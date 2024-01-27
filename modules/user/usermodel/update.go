package usermodel

import "bestHabit/common"

type UserUpdate struct {
	Phone    *string          `json:"phone" db:"phone"`
	Name     *string          `json:"name" db:"name"`
	Avatar   *common.Image    `json:"avatar" db:"avatar"`
	Settings *common.Settings `json:"settings" db:"settings"`
}
type UpdatePassword struct {
	NewPassword *string `json:"new_password"`
	Password    *string `json:"password" db:"password"`
}

type UpdateDeviceTokens struct {
	DeviceTokens *common.DvTokens `json:"device_tokens" db:"device_tokens"`
}

type ResetPassword struct {
	Password *string `json:"password" db:"password"`
}

func (UpdatePassword) TableName() string {
	return UserCreate{}.TableName()
}

func (UserUpdate) TableName() string {
	return UserCreate{}.TableName()
}
