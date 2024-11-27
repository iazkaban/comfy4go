package client

import (
	"math/rand/v2"
)

type InputType string

const (
	InputTypeRandom InputType = "rand_int64"
	InputTypeString InputType = "string"
	InputTypeInt64  InputType = "int64"
)

type Input struct {
	Type  InputType
	Value any
}

func (input *Input) GetValue() any {
	switch input.Type {
	case InputTypeInt64:
		return input.Value.(int64)
	case InputTypeString:
		return input.Value.(string)
	case InputTypeRandom:
		return rand.Int64()
	default:
		return 0
	}
}
