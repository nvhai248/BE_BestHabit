package usermodel

import "bestHabit/component/tokenprovider"

type UserLogin struct {
	Email    string `db:"email" json:"email" form:"email"`
	Password string `db:"password" json:"password" form:"password"`
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
