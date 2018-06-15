package main

import (
	"bufio"
	"log"
	"os"

	"github.com/k0kubun/pp"

	"github.com/Bo0km4n/praks"
)

func main() {
	p := praks.NewParser("json")

	fileName := "sample.json"
	fp, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(fp)
	scanner.Scan()

	s := p.TexToStruct(scanner.Text())
	pp.Println(s.GetValue("CHARSET"), s.GetValue("CLIENT_ID"), s.GetValue("NESTED"))
}
