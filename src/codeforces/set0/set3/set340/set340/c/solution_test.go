package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if !slices.Equal(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample(t *testing.T) {
	s := `3
2 3 5
`
	runSample(t, s, []int{22, 3})
}
