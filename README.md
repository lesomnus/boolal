# *bool*ean *al*gebra

Evaluate a boolean expression.

## Usage

```go
import ba "github.com/lesomnus/boolal"

func init(){
	data := map[string]bool{"t": true}
	expr := ba.ParseString("t & f | !(t | f)")

	ok := expr.Eval(data)
	// ok == false
}
```
