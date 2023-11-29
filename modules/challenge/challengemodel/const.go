package challengemodel

import (
	"bestHabit/common"
	"errors"
)

var (
	ErrNameNotBeBlank = common.NewCustomError(errors.New("Name not be blank!"),
		"Name of task not be blank!",
		"NameNotBeBlank")
	ErrStartDateNotBlank = common.NewCustomError(errors.New("StartDate not be blank!"),
		"StartDate of task not be blank!",
		"StartDateNotBeBlank")
	ErrEndDateNotBlank = common.NewCustomError(errors.New("EndDate not be blank!"),
		"EndDate of task not be blank!",
		"EndDateNotBeBlank")
)
