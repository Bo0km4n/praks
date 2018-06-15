package praks

import (
	"reflect"
	"strings"
)

// Struct is wrap struct
type Struct struct {
	Meta  reflect.Type
	Value reflect.Value
}

// GetFieldNames //
func (s *Struct) GetFieldNames() []string {
	names := make([]string, 0)
	for i := 0; i < s.Meta.NumField(); i++ {
		names = append(names, s.Meta.Field(i).Name)
	}
	return names
}

// GetValueTypeToString //
func (s *Struct) GetValueTypeToString(field string) string {
	v := s.Value.FieldByName(field)
	return v.Type().Name()
}

// GetFieldAndType //
func (s *Struct) GetFieldAndType() map[string]string {
	dict := map[string]string{}

	for i := 0; i < s.Meta.NumField(); i++ {
		name := s.Meta.Field(i).Name
		tName := s.Value.FieldByName(name).Type().Name()
		if tName == "" {
			tName = "map[string]interface {}"
		}
		dict[name] = tName
	}
	return dict
}

// SetValue //
func (s *Struct) SetValue(f string, v interface{}) {
	s.Value.FieldByName(f).Set(reflect.ValueOf(v))
}

// GetValue //
func (s *Struct) GetValue(f string) interface{} {
	f = strings.ToUpper(f)
	v := s.Value.FieldByName(f)
	return getCastedValue(v)
}

func getCastedValue(v reflect.Value) interface{} {
	typeName := v.Type().Name()

	switch typeName {
	case "string":
		return v.Interface().(string)
	case "int":
		return v.Interface().(int)
	case "int32":
		return v.Interface().(int32)
	case "int64":
		return v.Interface().(int64)
	case "float32":
		return v.Interface().(float32)
	case "float64":
		return v.Interface().(float64)
	case "bool":
		return v.Interface().(bool)
	default:
		return v.Interface()
	}
}
