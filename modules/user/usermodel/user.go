package usermodel

import (
	"bestHabit/common"
	"bestHabit/component/tokenprovider"
	"errors"
)

type User struct {
	common.SQLModel `json:",inline"`
	Email           *string          `json:"email" db:"email"`
	Phone           *string          `json:"phone" db:"phone"`
	Password        *string          `json:"-" db:"password"`
	Name            *string          `json:"name" db:"name"`
	FbID            *string          `json:"-" db:"fb_id"`
	GgID            *string          `json:"-" db:"gg_id"`
	Salt            *string          `json:"-" db:"salt"`
	Avatar          *common.Image    `json:"avatar" db:"avatar"`
	Level           int              `json:"level" db:"level"`
	Experience      int              `json:"experience" db:"experience"`
	Settings        *common.Settings `json:"settings" db:"settings"`
	Role            *string          `json:"role" db:"role"`
	HabitCount      int              `json:"habit_count" db:"habit_count"`
	TaskCount       int              `json:"task_count" db:"task_count"`
	ChallengeCount  int              `json:"challenge_count" db:"challenge_count"`
	Status          int              `json:"status" db:"status"`
}

func (User) TableName() string {
	return UserCreate{}.TableName()
}

func (user *User) Mask(isAdminOrOwner bool) {
	user.GenUID(common.DbTypeUser)
}

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
