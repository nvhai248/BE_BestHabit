package common

import "time"

func ParseStringToTimestamp(data string) (*time.Time, error) {
	result, err := time.Parse("2006-01-02 15:04:05", data)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func ParseStringToDate(data string) (*time.Time, error) {
	result, err := time.Parse("2006-01-02", data)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func ParseStringToTime(data string) (*time.Time, error) {
	result, err := time.Parse("15:04:05", data)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
