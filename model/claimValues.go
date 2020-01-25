// Copyright (c) 2020 Vorotynsky Maxim

package model

type (
	ClaimValue interface {
		GetValue() string
	}
	StringValue struct {
		Value string
	}
	HashValue struct {
		Hash []byte
	}
)

func (value StringValue) GetValue() string {
	return value.Value
}
func (value HashValue) GetValue() string {
	return string(value.Hash)
}
