// Code generated by queryx, DO NOT EDIT.

package queryx

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type String struct {
	Val   string
	Valid bool
	Set   bool
}

func NewString(v string) String {
	return String{Val: v, Valid: true, Set: true}
}

func NewNullableString(v *string) String {
	if v != nil {
		return NewString(*v)
	}
	return String{Set: true}
}

// Scan implements the Scanner interface.
func (s *String) Scan(value interface{}) error {
	ns := sql.NullString{String: s.Val}
	err := ns.Scan(value)
	s.Val, s.Valid = ns.String, ns.Valid
	return err
}

// Value implements the driver Valuer interface.
func (s String) Value() (driver.Value, error) {
	if !s.Valid {
		return nil, nil
	}
	return s.Val, nil
}

// MarshalJSON implements the json.Marshaler interface.
func (s String) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(s.Val)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s *String) UnmarshalJSON(data []byte) error {
	s.Set = true
	if string(data) == "null" {
		return nil
	}
	s.Valid = true
	if err := json.Unmarshal(data, &s.Val); err != nil {
		return err
	}
	return nil
}

// String implements the stringer interface.
func (s String) String() string {
	if !s.Valid {
		return "null"
	}
	return fmt.Sprintf(`"%s"`, s.Val)
}
