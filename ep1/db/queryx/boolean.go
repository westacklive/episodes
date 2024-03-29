// Code generated by queryx, DO NOT EDIT.

package queryx

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

type Boolean struct {
	Val   bool
	Valid bool
	Set   bool
}

func NewBoolean(v bool) Boolean {
	return Boolean{Val: v, Valid: true, Set: true}
}

func NewNullableBoolean(v *bool) Boolean {
	if v != nil {
		return NewBoolean(*v)
	}
	return Boolean{Set: true}
}

// Scan implements the Scanner interface.
func (b *Boolean) Scan(value interface{}) error {
	n := sql.NullBool{}
	err := n.Scan(value)
	if err != nil {
		return err
	}
	b.Val, b.Valid = n.Bool, n.Valid
	return nil
}

// Value implements the driver Valuer interface.
func (b Boolean) Value() (driver.Value, error) {
	if !b.Valid {
		return nil, nil
	}
	return b.Val, nil
}

// MarshalJSON implements the json.Marshaler interface.
func (b Boolean) MarshalJSON() ([]byte, error) {
	if !b.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(b.Val)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (b *Boolean) UnmarshalJSON(data []byte) error {
	b.Set = true
	if string(data) == "null" {
		return nil
	}
	b.Valid = true
	if err := json.Unmarshal(data, &b.Val); err != nil {
		return err
	}
	return nil
}

// String implements the stringer interface.
func (b Boolean) String() string {
	if !b.Valid {
		return "null"
	}
	if b.Val {
		return "true"
	}
	return "false"
}
