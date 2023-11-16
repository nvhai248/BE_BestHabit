package common

import "errors"

var (
	ErrorNoRows = errors.New("Error No Rows!")

	ErrEmailOrPasswordInvalid = NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"ErrEmailOrPasswordInvalid",
	)

	ErrEmailExisted = NewCustomError(
		errors.New("email has already been existed"),
		"email has already been existed",
		"ErrEmailExisted",
	)
)
