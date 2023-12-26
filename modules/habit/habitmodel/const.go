package habitmodel

import (
	"bestHabit/common"
	"errors"
)

var (
	ErrNameNotBeBlank = common.NewCustomError(errors.New("Name not be blank!"),
		"Name of task not be blank!",
		"NameNotBeBlank")

	ErrStartDateNotBeBlank = common.NewCustomError(errors.New("Start date not be blank!"),
		"Start date of task not be blank!",
		"StartDateNotBeBlank")

	ErrEndDateNotBeBlank = common.NewCustomError(errors.New("End date not be blank!"),
		"End date of task not be blank!",
		"EndDateNotBeBlank")

	ErrNameAlreadyUsed = common.NewCustomError(errors.New("Name already used by another task!"),
		"Name already used!",
		"NameAlreadyUsed")
)
