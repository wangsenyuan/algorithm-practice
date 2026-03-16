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
	s := `9 11
1
2
0 5 8
1
1
3
0 3 8
9
0 6 9
6
0 1 9
`
	expect := []int{2, 3, 2, 5}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 10
1
2
3
0 1 2
0 4 7
0 2 5
20
0 6 6
99
0 4 6
`
	expect := []int{1, 2, 2, 1, 3}
	runSample(t, s, expect)
}
