package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFibPositive(t *testing.T) {
	var b bytes.Buffer
	require.NoError(t, fib(&b, 7))
	require.Equal(t, b.String(), "1\n1\n2\n3\n5\n8\n13\n")
}

func TestFibNegative(t *testing.T) {
	var b bytes.Buffer
	require.NoError(t, fib(&b, -7))
	require.Equal(t, b.String(), "1\n-1\n2\n-3\n5\n-8\n13\n")
}
