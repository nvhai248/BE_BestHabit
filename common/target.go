package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Target struct {
	Times     int `json:"times"`
	TotalTime int `json:"total_time"`
}

func NewDefaultTarget() *Target {
	return &Target{Times: 0, TotalTime: 0}
}

func (j *Target) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		_ = errors.New(fmt.Sprint("Failed to unmarshal JSONB value: ", value))
	}

	var tg Target
	if err := json.Unmarshal(bytes, &tg); err != nil {
		return err
	}

	*j = tg
	return nil
}

func (j *Target) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}
