package taskmodel

import (
	"bestHabit/common"
	"errors"
)

var (
	ErrNameNotBeBlank = common.NewCustomError(errors.New("Name not be blank!"),
		"Name of task not be blank!",
		"NameNotBeBlank")

	ErrDeadlineNotBeBlank = common.NewCustomError(errors.New("Deadline not be blank!"),
		"Deadline of task not be blank!",
		"DeadlineNotBeBlank")

	ErrNameAlreadyUsed = common.NewCustomError(errors.New("Name already used by another task!"),
		"Name already used!",
		"NameAlreadyUsed")
)
