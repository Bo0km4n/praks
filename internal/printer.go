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
	if p.value.Len() == 0 {
		p.printf("%s{}", p.typeString())
		return
	}

	keys := p.value.MapKeys()
	keyType := keys[0].Type().String()
	valueType := p.value.MapIndex(keys[0]).Elem().Type().String()
	isMixed := false
	// check value types
	for i := 0; i < p.value.Len(); i++ {
		isMixed = (p.value.MapIndex(keys[i]).Elem().Type().String() != valueType)
	}

	if isMixed {
		valueType = "interface {}"
	}
	fmt.Println("ssssss", isMixed)
	p.printf("map < %s : %s >", keyType, valueType)

	// for i := 0; i < p.value.Len(); i++ {
	// 	value := p.value.MapIndex(keys[i])
	// 	p.printf("%s:\t%s", p.format(keys[i]), p.format(value))

	// 	if i < p.value.Len()-1 {
	// 		p.print(", ")
	// 	}
	// }
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

	v := p.value.Index(0)
	vp := NewPrinter(v)
	vp.Print()
	arrayType := fmt.Sprintf("array < %s >", vp.Flush())
	p.buffer.WriteString(arrayType)
	// for i := 0; i < p.value.Len(); i++ {
	// 	pp := NewPrinter(p.value.Index(i))
	// 	pp.Print()
	// 	p.buffer.WriteString(pp.Flush())
	// 	if i != p.value.Len()-1 && p.value.Len() > 1 {
	// 		p.buffer.WriteString(", ")
	// 	}
	// }
}

func (p *Printer) Flush() string {
	s := p.buffer.String()
	p.buffer = bytes.NewBufferString("")
	return s
}

func (p *Printer) format(object reflect.Value) string {
	pp := NewPrinter(object)
	pp.Print()
	return pp.Flush()
}

func (p *Printer) typeString() string {
	return p.value.Type().String()
}

func (p *Printer) printf(format string, args ...interface{}) {
	text := fmt.Sprintf(format, args...)
	p.print(text)
}

func (p *Printer) print(text string) {
	fmt.Fprint(p.buffer, text)
}
