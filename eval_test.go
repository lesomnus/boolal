package boolal_test

import (
	"testing"

	ba "github.com/lesomnus/boolal"
	"github.com/stretchr/testify/require"
)

func TestEval(t *testing.T) {
	data := map[string]bool{
		"t": true,
		"f": false,
	}

	tcs := []struct {
		desc     string
		input    *ba.Expr
		expected bool
	}{
		{
			desc:     "truth",
			input:    &ba.Expr{Lhs: ba.Truth()},
			expected: true,
		},
		{
			desc:     "falsity",
			input:    &ba.Expr{Lhs: ba.Falsity()},
			expected: false,
		},
		{
			desc:     "true",
			input:    &ba.Expr{Lhs: ba.Var("t")},
			expected: true,
		},
		{
			desc:     "false",
			input:    &ba.Expr{Lhs: ba.Var("f")},
			expected: false,
		},
		{
			desc:     "negation of true",
			input:    &ba.Expr{Lhs: ba.Not("t")},
			expected: false,
		},
		{
			desc:     "negation of false",
			input:    &ba.Expr{Lhs: ba.Not("f")},
			expected: true,
		},
		{
			desc:     "conjunctions of true",
			input:    ba.And("t", "t", "t"),
			expected: true,
		},
		{
			desc:     "conjunctions of false",
			input:    ba.And("f", "f", "f"),
			expected: false,
		},
		{
			desc:     "conjunctions of mixed booleans starts with true",
			input:    ba.And("t", "f", "t"),
			expected: false,
		},
		{
			desc:     "conjunctions of mixed booleans starts with false",
			input:    ba.And("f", "t", "f"),
			expected: false,
		},
		{
			desc:     "disjunctions of true",
			input:    ba.Or("t", "t", "t"),
			expected: true,
		},
		{
			desc:     "disjunctions of false",
			input:    ba.Or("f", "f", "f"),
			expected: false,
		},
		{
			desc:     "disjunctions of mixed booleans starts with true",
			input:    ba.Or("t", "f", "t"),
			expected: true,
		},
		{
			desc:     "disjunctions of mixed booleans starts with false",
			input:    ba.Or("f", "t", "f"),
			expected: true,
		},
		{
			desc:     "mixed junctions that ends with conjunction of false",
			input:    ba.And("t", "f").Or("t", "f").And("f"),
			expected: false,
		},
		{
			desc:     "mixed junctions that ends with disjunction of true",
			input:    ba.And("t", "f").Or("t", "f").And("t"),
			expected: true,
		},
		{
			desc:     "nested expressions starts with conjunction of false",
			input:    ba.And("f", ba.Or("t", ba.And("t", "t"))),
			expected: false,
		},
		{
			desc:     "nested expressions starts with disjunction of true",
			input:    ba.Or("t", ba.Or("f", ba.And("t", "f"))),
			expected: true,
		},
		{
			desc:     "nested truthy expressions",
			input:    ba.And("t", ba.Or("f", ba.And("t", "t"))),
			expected: true,
		},
		{
			desc:     "nested falsy expressions",
			input:    ba.And("t", ba.Or("f", ba.And("t", "f"))),
			expected: false,
		},
		{
			desc:     "negation of nested truthy expression",
			input:    ba.And("t", ba.Not(ba.And("t", "t"))),
			expected: false,
		},
		{
			desc:     "negation of nested falsy expression",
			input:    ba.And("t", ba.Not(ba.And("t", "f"))),
			expected: true,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			require := require.New(t)

			ok := tc.input.Eval(data)
			require.Equal(tc.expected, ok)
		})
	}

	t.Run("empty concat must be ignored", func(t *testing.T) {
		tcs := []struct {
			desc     string
			input    *ba.Expr
			expected bool
		}{
			{
				desc:     "true & ?",
				input:    &ba.Expr{Lhs: ba.Var("t"), Rhs: []ba.BinaryOp{{And: []ba.UnaryOp{}}}},
				expected: true,
			},
			{
				desc:     "false & ?",
				input:    &ba.Expr{Lhs: ba.Var("f"), Rhs: []ba.BinaryOp{{And: []ba.UnaryOp{}}}},
				expected: false,
			},
			{
				desc:     "true | ?",
				input:    &ba.Expr{Lhs: ba.Var("t"), Rhs: []ba.BinaryOp{{Or: []ba.UnaryOp{}}}},
				expected: true,
			},
			{
				desc:     "false | ?",
				input:    &ba.Expr{Lhs: ba.Var("f"), Rhs: []ba.BinaryOp{{Or: []ba.UnaryOp{}}}},
				expected: false,
			},
		}
		for _, tc := range tcs {
			t.Run(tc.desc, func(t *testing.T) {
				require := require.New(t)

				ok := tc.input.Eval(data)
				require.Equal(tc.expected, ok)
			})
		}
	})
}

func TestEvalWithFallbackValue(t *testing.T) {
	require := require.New(t)

	data := map[string]bool{"t": true}
	expr := ba.And("t", "x")

	{
		ok := expr.Eval(data, ba.WithFallbackAs(false))
		require.False(ok)
	}

	{
		ok := expr.Eval(data, ba.WithFallbackAs(true))
		require.True(ok)
	}
}
