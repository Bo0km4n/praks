package praks

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type jsonParser struct {
	timeFormat string
}

func (j *jsonParser) TexToStruct(text string) *Struct {
	var s Struct
	item, err := j.parseText([]byte(text))
	if err != nil {
		return &s
	}
	rt := j.buildReflectType(item)
	s.Meta = rt
	s.Value = reflect.New(rt).Elem()
	j.setDefaultValue(&s, item)
	return &s
}

func (j *jsonParser) SetTimeFormat(format string) {
	j.timeFormat = format
}

func (j *jsonParser) parseText(body []byte) (interface{}, error) {
	var item interface{}

	err := json.Unmarshal(body, &item)
	return item, err
}

func (j *jsonParser) setDefaultValue(s *Struct, item interface{}) {
	dict := item.(map[string]interface{})
	for k, v := range dict {
		s.SetValue(strings.ToUpper(k), v)
	}
}

func (j *jsonParser) buildReflectType(item interface{}) reflect.Type {
	dict := item.(map[string]interface{})
	fields := make([]reflect.StructField, 0)
	for i, v := range dict {
		fmt.Println(reflect.TypeOf(v))
		f := reflect.StructField{Name: strings.ToUpper(i), Type: reflect.TypeOf(v)}
		jsonTag := fmt.Sprintf("json:\"%s\"", i)
		tag := reflect.ValueOf(&f.Tag).Elem()
		tag.SetString(jsonTag)
		fields = append(fields, f)
	}
	return reflect.StructOf(fields)
}
