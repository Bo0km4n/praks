package internal

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type Printer struct {
	buffer *bytes.Buffer
	value  reflect.Value
}

func NewPrinter(object reflect.Value) *Printer {
	return &Printer{
		buffer: bytes.NewBufferString(""),
		value:  object,
	}
}

func (p *Printer) Print() {
	switch p.value.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64,
		reflect.Float32, reflect.Float64, reflect.Bool,
		reflect.String:
		p.buffer.WriteString(fmt.Sprintf("%v", p.value.Kind()))
	case reflect.Map:
		p.printMap()
	case reflect.Struct:
		p.printStruct()
	case reflect.Slice, reflect.Array:
		p.printSlice()
	default:
		p.printInterface(p.value)
	}
}

func (p *Printer) printMap() {
	m := p.value
	p.buffer.WriteString("map{ ")
	keys := m.MapKeys()
	for i := 0; i < m.Len(); i++ {
		v := m.MapIndex(keys[i])
		p.value = v
		p.Print()
		if i != m.Len()-1 && m.Len() > 1 {
			p.buffer.WriteString(", ")
		}
	}
	p.buffer.WriteString(" }")
}

func (p *Printer) printInterface(v reflect.Value) {
	e := v.Elem()
	if e.Kind() == reflect.Invalid {
		p.buffer.WriteString("nil")
	} else if e.IsValid() {
		p.value = e
		p.Print()
	} else {
		t := p.value.Type().String()
		p.buffer.WriteString(t + "(nil)")
	}
}

func (p *Printer) printStruct() {
	if p.value.Type().String() == "time.Time" {
		p.printTime()
		return
	}
}

func (p *Printer) printTime() {
	if !p.value.CanInterface() {
		p.buffer.WriteString("(unexported time.Time)")
		return
	}
	tm := p.value.Interface().(time.Time)
	t := fmt.Sprintf(
		"%s-%s-%sT%s:%s:%s",
		strconv.Itoa(tm.Year()),
		fmt.Sprintf("%02d", tm.Month()),
		fmt.Sprintf("%02d", tm.Day()),
		fmt.Sprintf("%02d", tm.Hour()),
		fmt.Sprintf("%02d", tm.Minute()),
		fmt.Sprintf("%02d", tm.Second()),
	)
	p.buffer.WriteString(t)
}

func (p *Printer) printSlice() {
	if p.value.Len() == 0 {
		st := fmt.Sprintf("%s[]", p.value.Type().String())
		p.buffer.WriteString(st)
		return
	}

	p.buffer.WriteString("array[ ")
	for i := 0; i < p.value.Len(); i++ {
		pp := NewPrinter(p.value.Index(i))
		pp.Print()
		p.buffer.WriteString(pp.Flush())
		if i != p.value.Len()-1 && p.value.Len() > 1 {
			p.buffer.WriteString(", ")
		}
	}
	p.buffer.WriteString(" ]")
}

func (p *Printer) Flush() string {
	s := p.buffer.String()
	p.buffer = bytes.NewBufferString("")
	return s
}
