package boolal_test

import (
	"testing"

	ba "github.com/lesomnus/boolal"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	tcs := []struct {
		desc     string
		input    string
		expected *ba.Expr
	}{
		{
			desc:     "No op",
			input:    "x",
			expected: &ba.Expr{Lhs: ba.Var("x")},
		},
		{
			desc:     "Negation",
			input:    "!x",
			expected: &ba.Expr{Lhs: ba.Not("x")},
		},
		{
			desc:     "Conjunction",
			input:    "x & y",
			expected: ba.And("x", "y"),
		},
		{
			desc:     "Disjunction",
			input:    "x | y",
			expected: ba.Or("x", "y"),
		},
		{
			desc:     "Conjunctions",
			input:    "x & y & z",
			expected: ba.And("x", "y", "z"),
		},
		{
			desc:     "Disjunctions",
			input:    "x | y | z",
			expected: ba.Or("x", "y", "z"),
		},
		{
			desc:     "Mixed junctions",
			input:    "x & !y | z",
			expected: ba.And("x", ba.Not("y")).Or("z"),
		},
		{
			desc:     "sub-expression on right side",
			input:    "x & (y | z)",
			expected: ba.And("x", ba.Or("y", "z")),
		},
		{
			desc:     "sub-expression on left side",
			input:    "(x & y) | z",
			expected: ba.Or(ba.And("x", "y"), "z"),
		},
		{
			desc:     "nested sub-expressions",
			input:    "x & (y | (z & a))",
			expected: ba.And("x", ba.Or("y", ba.And("z", "a"))),
		},
		{
			desc:     "sub-expressions with unary op",
			input:    "x & !(y | z)",
			expected: ba.And("x", ba.Not(ba.Or("y", "z"))),
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			require := require.New(t)

			al, err := ba.ParseString(tc.input)
			require.NoError(err)
			require.Equal(tc.expected, al)
		})
	}
}
