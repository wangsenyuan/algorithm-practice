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
	s := `4 1
1 1 1 2
1 1 1 1
4
1 2
1 3
1 4
3 4
`
	expect := []int{2, 3, 4, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 0
1 2 1 2
0 0 0 0
1
1 4
`
	expect := []int{10}
	runSample(t, s, expect)
}
