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
	s := `2 2
1 3
1
3
`
	expect := []int{1, 0}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 3
0 1 5 6
1
2
4
`
	expect := []int{2, 0, 0}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 4
0 1 5 6 7
1
1
4
5
`
	expect := []int{2, 2, 0, 2}
	runSample(t, s, expect)
}
