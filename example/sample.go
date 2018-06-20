package main

import (
	"bufio"
	"log"
	"os"

	"github.com/k0kubun/pp"

	"github.com/Bo0km4n/praks"
)

func main() {
	p, err := praks.NewParser("json")

	if err != nil {
		log.Fatal(err)
	}

	fileName := "sample.json"
	fp, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(fp)
	scanner.Scan()

	s := p.TexToStruct(scanner.Text())
	pp.Println(s.GetValue("CHARSET"))
	pp.Println(s.GetValue("NEST_NEST"))
	pp.Println(s.GetValue("time"))
	pp.Println(s.GetFieldAndType())
	pp.Println(s.Value.Interface())
}
