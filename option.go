// Package opt provides a Value type that can be used to represent values that
// may be present or absent. JSON marshaling and unmarshaling of those value
// types are simplified, if the fields are tagged with `json:",omitzero"`.
package opt

import "encoding/json"

// Value represents an optional Value. The value V can only be interpreted if Ok
// is true. Note that the zero value is not a valid entry.
type Value[T any] struct {
	V  T
	Ok bool
}

// Returns a valid Value.
func Val[T any](v T) Value[T] { return Value[T]{V: v, Ok: true} }

// NoVal returns an invalid Value.
func NoVal[T any]() Value[T] {
	var zero T
	return Value[T]{V: zero, Ok: false}
}

// MarshalJSON implements the json.Marshaler interface. If the Value is valid,
// it marshals the value V. If the Value is invalid, it marshals to null.
func (o Value[T]) MarshalJSON() ([]byte, error) {
	if o.Ok {
		return json.Marshal(o.V)
	}
	return []byte("null"), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface. If the input is null,
// it sets the Value to invalid. Otherwise, it tries to unmarshal the input into
// the value V and sets the Value to valid if successful.
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
