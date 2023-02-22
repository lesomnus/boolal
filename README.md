# *bool*ean *al*gebra

[![test](https://github.com/lesomnus/boolal/actions/workflows/test.yaml/badge.svg)](https://github.com/lesomnus/boolal/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/lesomnus/boolal)](https://goreportcard.com/report/github.com/lesomnus/boolal)
[![codecov](https://codecov.io/gh/lesomnus/boolal/branch/main/graph/badge.svg?token=9JMg9rhj2w)](https://codecov.io/gh/lesomnus/boolal)

Evaluate a boolean expression.

## Usage

```go
import ba "github.com/lesomnus/boolal"

func Expression() {
	data := map[string]bool{"t": true}
	expr, err := ba.ParseString("t & f | !(t | f)")
	if err != nil {
		panic(err)
	}

	ok := expr.Eval(data)
	// ok == false
}

func Manipulation() {
	data := map[string]bool{"t": true}
	expr := ba.And("t", "f").Or(ba.Not(ba.Or("t", "f")))

	ok := expr.Eval(data)
	// ok == false
}
```
