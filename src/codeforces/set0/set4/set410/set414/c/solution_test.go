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
	s := `2
2 1 4 3
4
1 2 0 2
`
	expect := []int{0, 6, 6, 0}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1
1 2
3
0 1 1
`
	expect := []int{0, 1, 0}
	runSample(t, s, expect)
}
