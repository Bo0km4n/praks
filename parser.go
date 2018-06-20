package praks

import (
	"errors"
	"reflect"
	"time"
)

// Parser interface //
type Parser interface {
	TexToStruct(text string) *Struct
	GetTimeFormat() string
	SetTagFormat(string)
}

var defaultTimeFormat = "2006-01-02T15:04:05"
var defaultTagFormat = "json"

// NewParser get json or csv parser
func NewParser(t string) (Parser, error) {
	switch t {
	case "json":
		return newJSONParser(defaultTimeFormat, defaultTagFormat), nil
	// TODO
	// Add csv case
	default:
		return nil, errors.New("not found parser")
	}
}

func getCastedValue(p Parser, v reflect.Value) interface{} {
	typeName := v.Type().Name()

	switch typeName {
	case "string":
		return castString(p, v.Interface().(string))
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

func castString(p Parser, v string) interface{} {
	t, err := time.Parse(p.GetTimeFormat(), v)
	if err != nil {
		return v
	}
	return t
}
