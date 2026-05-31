// Package opt provides a Value type that can be used to represent values that
// may be present or absent. JSON marshaling and unmarshaling of those value
// types are simplified, if the fields are tagged with `json:",omitzero"`.
package opt

import "encoding/json"

type Value[T any] struct {
	V  T
	Ok bool
}

func Val[T any](v T) Value[T] { return Value[T]{V: v, Ok: true} }

func NoVal[T any]() Value[T] {
	var zero T
	return Value[T]{V: zero, Ok: false}
}

func (o Value[T]) MarshalJSON() ([]byte, error) {
	if o.Ok {
		return json.Marshal(o.V)
	}
	return []byte("null"), nil
}

func (o *Value[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.Ok = false
		var zero T
		o.V = zero
		return nil
	}
	err := json.Unmarshal(data, &o.V)
	if err != nil {
		return err
	}
	o.Ok = true
	return nil
}
