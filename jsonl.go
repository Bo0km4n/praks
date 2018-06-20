package praks

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type jsonParser struct {
	timeFormat string
	tagFormat  string
}

func newJSONParser(timeF, tagF string) *jsonParser {
	return &jsonParser{
		timeFormat: timeF,
		tagFormat:  tagF,
	}
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
	s.parser = j
	j.setDefaultValue(&s, item)
	return &s
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
		f := reflect.StructField{Name: strings.ToUpper(i), Type: reflect.TypeOf(getCastedValue(j, reflect.ValueOf(v)))}
		jsonTag := fmt.Sprintf("%s:\"%s\"", j.tagFormat, i)
		tag := reflect.ValueOf(&f.Tag).Elem()
		tag.SetString(jsonTag)
		fields = append(fields, f)
	}
	return reflect.StructOf(fields)
}

func (j *jsonParser) GetTimeFormat() string {
	return j.timeFormat
}

func (j *jsonParser) SetTimeFormat(f string) {
	j.timeFormat = f
}

func (j *jsonParser) SetTagFormat(f string) {
	j.tagFormat = f
}
