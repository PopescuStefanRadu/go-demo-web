package types

import "encoding/json"

type Option[T any] struct {
	IsPresent bool
	Val       *T
}

func (o *Option[T]) UnmarshalJSON(data []byte) error {
	o.IsPresent = true
	if string(data) == "null" {
		return nil
	}
	return json.Unmarshal(data, &o.Val)
}

func NewOptionNillable[T any](val *T) Option[T] {
	return Option[T]{
		IsPresent: true,
		Val:       val,
	}
}

func NewOption[T any](val T) Option[T] {
	return Option[T]{
		IsPresent: true,
		Val:       &val,
	}
}
