package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type DvToken struct {
	DeviceToken string `json:"device_token"`
}

func (j *DvToken) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		_ = errors.New(fmt.Sprint("Failed to unmarshal JSONB value: ", value))
	}

	var DvToken DvToken
	if err := json.Unmarshal(bytes, &DvToken); err != nil {
		return err
	}

	*j = DvToken
	return nil
}

func (j *DvToken) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}

type DvTokens []DvToken

func (j *DvTokens) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		_ = errors.New(fmt.Sprint("Failed to unmarshal JSONB value: ", value))
	}

	var DvTokens []DvToken
	if err := json.Unmarshal(bytes, &DvTokens); err != nil {
		return err
	}

	*j = DvTokens
	return nil
}

func (j *DvTokens) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}

func (j *DvTokens) Init() {
	*j = make(DvTokens, 0)
}

func (j *DvTokens) AddNewDvToken(data DvToken) {
	*j = append(*j, data)
}
