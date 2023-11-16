package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Settings struct {
	Theme    string `json:"theme"`
	Language string `json:"language"`
}

func NewDefaultSettings() *Settings {
	return &Settings{Theme: "light", Language: "en"}
}

func (j *Settings) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		errors.New(fmt.Sprint("Failed to unmarshal JSONB value: ", value))
	}

	var img Settings
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img
	return nil
}

func (j *Settings) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}
