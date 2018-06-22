package main

import (
	"github.com/k0kubun/pp"

	"github.com/Bo0km4n/praks"
)

func main() {
	j := `{
		"team": "github.com",
		"fqdn": "google.com",
		"version": "0.9.0",
		"charset": "utf-8",
		"custom": {
		  "nest_custom": {
			  "col1": "ooo",
			  "col2": 100
		  },
		  "multiple-condition-1": 1,
		  "multiple-condition-2": 2
		},
		"global_id": "8e23594e-2927-4a44-a995-4821caa9f5db",
		"host": "localhost:9000",
		"language": "ja",
		"title": "",
		"array": [
			{"obj1": "hello"},
			{"obj2": "pokemon"}
		],
		"array_complex": [
			"hello",
			"hello" 
		]
	}`
	p, err := praks.NewParser("json")
	if err != nil {
		return
	}
	s := p.TexToStruct(j)
	pp.Println(s.GetFieldAndType())
	pp.Println(s.Value.Interface())
}
