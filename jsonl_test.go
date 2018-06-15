package praks

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONParse(t *testing.T) {
	p := jsonParser{}

	fileName := "example/sample.json"
	fp, err := os.Open(fileName)
	if err != nil {
		t.Fatal(err)
	}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		item, err := p.parseText([]byte(scanner.Text()))
		assert.NotNil(t, item)
		assert.Equal(t, nil, err)
	}
}

func TestBuildReflectType(t *testing.T) {
	p := jsonParser{}

	fileName := "example/sample.json"
	fp, err := os.Open(fileName)
	if err != nil {
		t.Fatal(err)
	}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		item, err := p.parseText([]byte(scanner.Text()))
		assert.Equal(t, nil, err)

		rt := p.buildReflectType(item)
		t.Log(rt)
	}
}

func TestSetStructFieldType(t *testing.T) {
	p := jsonParser{}

	fileName := "example/sample.json"
	fp, err := os.Open(fileName)
	if err != nil {
		t.Fatal(err)
	}
	scanner := bufio.NewScanner(fp)
	scanner.Scan()
	s := p.TexToStruct(scanner.Text())

	assert.NotNil(t, s.Meta)
	assert.NotNil(t, s.Value)
}
