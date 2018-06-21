package praks

import (
	"reflect"
	"strings"

	"github.com/Bo0km4n/praks/internal"
)

// Struct is wrap struct
type Struct struct {
	Meta   reflect.Type
	Value  reflect.Value
	parser Parser
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
		p := internal.NewPrinter(s.Value.FieldByName(name))
		p.Print()
		dict[name] = p.Flush()
	}
	return dict
}

// SetValue //
func (s *Struct) SetValue(f string, v interface{}) {
	cv := getCastedValue(s.parser, reflect.ValueOf(v))
	s.Value.FieldByName(f).Set(reflect.ValueOf(cv))
}

// GetValue //
func (s *Struct) GetValue(f string) interface{} {
	f = strings.ToUpper(f)
	v := s.Value.FieldByName(f)
	return getCastedValue(s.parser, v)
}
