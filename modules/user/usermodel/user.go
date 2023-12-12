package usermodel

import (
	"bestHabit/common"
	"bestHabit/component/tokenprovider"
	"errors"
)

type UserLogin struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

func (UserLogin) TableName() string {
	return UserCreate{}.TableName()
}

type Account struct {
	AccessToken *tokenprovider.Token `json:"access_token"`
}

func NewAccount(at *tokenprovider.Token) *Account {
	return &Account{AccessToken: at}
}

var (
	ErrNameCannotBeEmpty = common.NewCustomError(
		errors.New("Name cannot be empty!"),
		"Name cannot be empty!",
		"NameCannotBeEmpty")

	ErrEmailCannotBeEmpty = common.NewCustomError(
		errors.New("Email cannot be empty!"),
		"Email cannot be empty!",
		"EmailCannotBeEmpty")
)
