package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type CompleteDate struct {
	Date      string `json:"date"`
	Times     int    `json:"times"`
	TotalTime int    `json:"total_time"`
}

func (j *CompleteDate) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		_ = errors.New(fmt.Sprint("Failed to unmarshal JSONB value: ", value))
	}

	var day CompleteDate
	if err := json.Unmarshal(bytes, &day); err != nil {
		return err
	}

	*j = day
	return nil
}

func (j *CompleteDate) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}

type CompleteDates []CompleteDate

func (j *CompleteDates) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		_ = errors.New(fmt.Sprint("Failed to unmarshal JSONB value: ", value))
	}

	var days []CompleteDate
	if err := json.Unmarshal(bytes, &days); err != nil {
		return err
	}

	*j = days
	return nil
}

func (j *CompleteDates) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}

func (j *CompleteDates) AddDate(date CompleteDate) {
	*j = append(*j, date)
}

func (j *CompleteDates) Init() {
	*j = make(CompleteDates, 0)
}
