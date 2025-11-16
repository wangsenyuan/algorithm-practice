package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)

	if !slices.Equal(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
1 2 3
`
	expect := []int{1, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
4 2 3 1 6
`
	expect := []int{4, 2, 2, 4}
	runSample(t, s, expect)
}
