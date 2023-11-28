package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Day struct {
	Weekday string `json:"weekday"`
}

func (j *Day) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		_ = errors.New(fmt.Sprint("Failed to unmarshal JSONB value: ", value))
	}

	var day Day
	if err := json.Unmarshal(bytes, &day); err != nil {
		return err
	}

	*j = day
	return nil
}

func (j *Day) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}

type Days []Day

func (j *Days) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		_ = errors.New(fmt.Sprint("Failed to unmarshal JSONB value: ", value))
	}

	var days []Day
	if err := json.Unmarshal(bytes, &days); err != nil {
		return err
	}

	*j = days
	return nil
}

func (j *Days) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}

func (j *Days) Init() {
	*j = make(Days, 0)
}
