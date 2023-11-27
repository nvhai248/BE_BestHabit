package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Date struct {
	Date string `json:"date"`
}

func (j *Date) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		_ = errors.New(fmt.Sprint("Failed to unmarshal JSONB value: ", value))
	}

	var day Date
	if err := json.Unmarshal(bytes, &day); err != nil {
		return err
	}

	*j = day
	return nil
}

func (j *Date) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}

type Dates []Date

func (j *Dates) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		_ = errors.New(fmt.Sprint("Failed to unmarshal JSONB value: ", value))
	}

	var days []Date
	if err := json.Unmarshal(bytes, &days); err != nil {
		return err
	}

	*j = days
	return nil
}

func (j *Dates) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}

func (j *Dates) AddDate(date Date) {
	*j = append(*j, date)
}

func (j *Dates) Init() {
	*j = make(Dates, 0)
}
