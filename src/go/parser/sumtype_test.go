package parser

import (
	"go/ast"
	"go/token"
	"testing"
)

const code = `
package p

import fmt "fmt"

type AB := A | B | C | D
type IS := struct{} | interface{} | []int | map[string]string | chan struct{}

func main() {}
`

func TestSumTypeParse(t *testing.T) {
	fset := token.NewFileSet()

	f, err := ParseFile(fset, "s.go", code, ParseComments)
	if err != nil {
		t.Fatal(err)
	}

	ast.Inspect(f, func(node ast.Node) bool {
		if v, ok := node.(*ast.SumType); ok {
			t.Log(v.Variants)
		}
		return true
	})
}
