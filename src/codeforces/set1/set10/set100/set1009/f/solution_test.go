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
	s := `4
1 2
2 3
3 4
`
	expect := []int{0, 0, 0, 0}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
1 2
1 3
1 4
`
	expect := []int{1, 0, 0, 0}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
1 2
2 3
2 4
`
	expect := []int{2, 1, 0, 0}
	runSample(t, s, expect)
}
