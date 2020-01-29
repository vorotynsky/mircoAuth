// Copyright (c) 2020 Vorotynsky Maxim

package model

import (
	"encoding/json"
	"errors"
	"time"
)

type Duration time.Duration

func (d *Duration) UnmarshalJSON(b []byte) error {
	var value interface{}
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}
	switch value := value.(type) {
	case string:
		duration, err := time.ParseDuration(value)
		if err != nil {
			return err
		}
		*d = Duration(duration)
	default:
		return errors.New("invalid duration")
	}
	return nil
}
