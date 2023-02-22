package boolal_test

import (
	"testing"

	ba "github.com/lesomnus/boolal"
	"github.com/stretchr/testify/require"
)

func TestExpression(t *testing.T) {
	require := require.New(t)

	data := map[string]bool{"t": true}
	expr, err := ba.ParseString("t & f | !(t | f)")
	require.NoError(err)

	ok := expr.Eval(data)
	require.False(ok)
}

func TestManipulation(t *testing.T) {
	require := require.New(t)

	data := map[string]bool{"t": true}
	expr := ba.And("t", "f").Or(ba.Not(ba.Or("t", "f")))

	ok := expr.Eval(data)
	require.False(ok)
}
