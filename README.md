# Praks

praks is a library for dynamically generating golang structure code from json and csv text data.

Supported text style:
- json
- csv <= will

## Installation

Use the `go` command:

```
$ go get github.com/Bo0km4n/praks
```

## Example

```go
package main

import (
	"bufio"
	"log"
	"os"

	"github.com/Bo0km4n/praks"
)

func main() {
    jsonBody := `{"charset":"utf-8","client_id":"dfd0ba49-d8ad-4c53-86e5-b05018ae5b90", "nested": {"col0": 1, "col1": "hello"}}`

	p := praks.NewParser("json")
	s := p.TexToStruct(jsonBody)
	fmt.Println(s.GetValue("CHARSET"), s.GetValue("CLIENT_ID"), s.GetValue("NESTED"))
}
```
