package praks

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFieldNames(t *testing.T) {

	tc := []struct {
		Parser   string
		FilePath string
		Expect   int
	}{
		{
			Parser:   "json",
			FilePath: "example/sample.json",
			Expect:   5,
		},
	}

	for _, c := range tc {
		p, err := NewParser(c.Parser)
		if err != nil {
			t.Fatal(err)
		}
		fp, err := os.Open(c.FilePath)
		if err != nil {
			t.Fatal(err)
		}

		scanner := bufio.NewScanner(fp)
		scanner.Scan()
		s := p.TexToStruct(scanner.Text())
		f := s.GetFieldNames()
		assert.Equal(t, c.Expect, len(f))
	}
}

func TestGetValueTypeToStringJSON(t *testing.T) {
	tc := []struct {
		Parser   string
		FilePath string
		Expect   string
	}{
		{
			Parser:   "json",
			FilePath: "example/sample.json",
			Expect:   "string",
		},
	}

	for _, c := range tc {
		p, err := NewParser(c.Parser)
		if err != nil {
			t.Fatal(err)
		}
		fp, err := os.Open(c.FilePath)
		if err != nil {
			t.Fatal(err)
		}

		scanner := bufio.NewScanner(fp)
		scanner.Scan()
		s := p.TexToStruct(scanner.Text())
		assert.Equal(t, c.Expect, s.GetValueTypeToString("CHARSET"))
	}
}

func TestGetFieldAndType(t *testing.T) {
	tc := []struct {
		Parser   string
		FilePath string
		Keys     []string
	}{
		{
			Parser:   "json",
			FilePath: "example/sample.json",
			Keys:     []string{"CHARSET", "CLIENT_ID"},
		},
	}

	for _, c := range tc {
		p, err := NewParser(c.Parser)
		if err != nil {
			t.Fatal(err)
		}
		fp, err := os.Open(c.FilePath)
		if err != nil {
			t.Fatal(err)
		}

		scanner := bufio.NewScanner(fp)
		scanner.Scan()
		s := p.TexToStruct(scanner.Text())
		d := s.GetFieldAndType()

		for _, k := range c.Keys {
			_, ok := d[k]
			assert.True(t, ok)
		}
	}
}
