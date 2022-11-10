package boolal_test

import (
	"testing"

	ba "github.com/lesomnus/boolal"
	"github.com/stretchr/testify/require"
)

func TestManipInvalidVariable(t *testing.T) {
	require.Panics(t, func() { ba.Not(42) }, "invalid")
}

func TestManip(t *testing.T) {
	tcs := []struct {
		desc     string
		input    *ba.Expr
		expected *ba.Expr
	}{
		{
			desc:  "and",
			input: ba.And("x", ba.Not("y"), "z"),
			expected: &ba.Expr{
				Lhs: ba.UnaryOp{Var: "x"},
				Rhs: []ba.BinaryOp{{And: []ba.UnaryOp{
					{Name: "!", Var: "y"},
					{Var: "z"},
				}}},
			},
		},
		{
			desc:  "or",
			input: ba.Or("x", ba.Not("y"), "z"),
			expected: &ba.Expr{
				Lhs: ba.UnaryOp{Var: "x"},
				Rhs: []ba.BinaryOp{{Or: []ba.UnaryOp{
					{Name: "!", Var: "y"},
					{Var: "z"},
				}}},
			},
		},
		{
			desc:  "concat",
			input: ba.And("x", ba.Not("y"), "z").Or(ba.Not("a"), "b").And("c", ba.Not("d")),
			expected: &ba.Expr{
				Lhs: ba.UnaryOp{Var: "x"},
				Rhs: []ba.BinaryOp{
					{And: []ba.UnaryOp{
						{Name: "!", Var: "y"},
						{Var: "z"},
					}},
					{Or: []ba.UnaryOp{
						{Name: "!", Var: "a"},
						{Var: "b"},
					}},
					{And: []ba.UnaryOp{
						{Var: "c"},
						{Name: "!", Var: "d"},
					}},
				},
			},
		},
		{
			desc:  "nested",
			input: ba.And("x", ba.Or("y", "z"), "a"),
			expected: &ba.Expr{
				Lhs: ba.UnaryOp{Var: "x"},
				Rhs: []ba.BinaryOp{{And: []ba.UnaryOp{
					{Al: &ba.Expr{
						Lhs: ba.UnaryOp{Var: "y"},
						Rhs: []ba.BinaryOp{{Or: []ba.UnaryOp{
							{Var: "z"},
						}}},
					}},
					{Var: "a"},
				}}},
			},
		},
		{
			desc:  "nested with not",
			input: ba.And("x", ba.Not(ba.Or("y", "z")), "a"),
			expected: &ba.Expr{
				Lhs: ba.UnaryOp{Var: "x"},
				Rhs: []ba.BinaryOp{{And: []ba.UnaryOp{
					{Name: "!", Al: &ba.Expr{
						Lhs: ba.UnaryOp{Var: "y"},
						Rhs: []ba.BinaryOp{{Or: []ba.UnaryOp{
							{Var: "z"},
						}}},
					}},
					{Var: "a"},
				}}},
			},
		},
		{
			desc:  "not not",
			input: ba.And("x", ba.Not(ba.Not("y")), "z"),
			expected: &ba.Expr{
				Lhs: ba.UnaryOp{Var: "x"},
				Rhs: []ba.BinaryOp{{And: []ba.UnaryOp{
					{Var: "y"},
					{Var: "z"},
				}}},
			},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tc.expected, tc.input)
		})
	}
}

func TestSingleVarExprConcat(t *testing.T) {
	require := require.New(t)

	{
		expr := &ba.Expr{Lhs: ba.Var("x")}
		expr.And("y")

		require.Equal(ba.And("x", "y"), expr)
	}

	{
		expr := &ba.Expr{Lhs: ba.Var("x")}
		expr.Or("y")

		require.Equal(ba.Or("x", "y"), expr)
	}

	{
		expr := &ba.Expr{Lhs: ba.Var("x")}
		expr.Or(&ba.Expr{Lhs: ba.Var("y")})

		require.Equal(ba.Or("x", "y"), expr)
	}
}
