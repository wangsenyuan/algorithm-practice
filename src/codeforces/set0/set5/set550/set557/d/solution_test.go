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
	s := `4 4
1 2
1 3
4 2
4 3
`
	expect := []int{1, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3
1 2
2 3
3 1
`
	expect := []int{0, 1}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 0`
	expect := []int{3, 1}
	runSample(t, s, expect)
}
