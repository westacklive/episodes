// Code generated by queryx, DO NOT EDIT.

package queryx

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

func ScanOne(rows *sql.Rows, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return fmt.Errorf("must pass pointer")
	}

	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	typ := rv.Type().Elem()

	// column -> index
	names := make(map[string]int, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		names[columnName(f)] = i
	}

	if !rows.Next() {
		return sql.ErrNoRows
	}

	indexes := make(map[int]int, typ.NumField())
	for i, c := range columns {
		name := strings.ToLower(strings.Split(c, "(")[0])
		index, ok := names[name]
		if !ok {
			return fmt.Errorf("name not found")
		}
		indexes[i] = index
	}

	values := make([]interface{}, len(columns))
	for i := range columns {
		t := typ.Field(indexes[i]).Type
		values[i] = reflect.New(t).Interface()

	}

	// scan into interfaces
	if err := rows.Scan(values...); err != nil {
		return err
	}

	for i, v := range values {
		reflect.Indirect(rv).Field(indexes[i]).Set(reflect.Indirect(reflect.ValueOf(v)))
	}

	if rows.Next() {
		return fmt.Errorf("more than one row")
	}

	return nil
}

func ScanSlice(rows *sql.Rows, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return fmt.Errorf("must pass pointer")
	}

	rv = reflect.Indirect(rv)
	if k := rv.Kind(); k != reflect.Slice {
		return fmt.Errorf("must pass slice")
	}

	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	typ := rv.Type().Elem()

	names := make(map[string]int, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		names[columnName(f)] = i
	}

	indexes := make(map[int]int, typ.NumField())
	for i, c := range columns {
		name := strings.ToLower(strings.Split(c, "(")[0])
		index, ok := names[name]
		if !ok {
			return fmt.Errorf("name %+v not found", name)
		}
		indexes[i] = index
	}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		for i := range columns {
			t := typ.Field(indexes[i]).Type
			values[i] = reflect.New(t).Interface()
		}

		// scan into interfaces
		if err := rows.Scan(values...); err != nil {
			return err
		}

		// convert to reflect.Value
		e := reflect.New(typ).Elem()
		for i, v := range values {
			fv := e.Field(indexes[i])
			fv.Set(reflect.Indirect(reflect.ValueOf(v)))
		}

		vv := reflect.Append(rv, e)
		rv.Set(vv)
	}

	return nil
}

func columnName(f reflect.StructField) string {
	name := strings.ToLower(f.Name)
	if tag, ok := f.Tag.Lookup("db"); ok {
		name = tag
	}
	return name
}
