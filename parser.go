package boolal

import "github.com/alecthomas/participle/v2"

type Expr struct {
	Lhs UnaryOp    `parser:"  @@"`
	Rhs []BinaryOp `parser:"@@*"`
}

type BinaryOp struct {
	And []UnaryOp `parser:"  ( '&' @@ )+"`
	Or  []UnaryOp `parser:"| ( '|' @@ )+"`
}

type UnaryOp struct {
	Name string `parser:"@( '!' )?"`
	Var  string `parser:"(   @Ident"`
	Al   *Expr  `parser:"  | '(' @@ ')' )"`
}

var bl_parser = participle.MustBuild[Expr]()

func ParseString(expr string) (*Expr, error) {
	return bl_parser.ParseString("", expr)
}
