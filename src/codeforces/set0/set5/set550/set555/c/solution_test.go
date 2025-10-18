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
	s := `6 5
3 4 U
6 1 L
2 5 L
1 6 U
4 3 U
`
	expect := []int{4, 3, 2, 1, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 6
2 9 U
10 1 U
1 10 U
8 3 L
10 1 L
6 5 U
`
	expect := []int{9, 1, 10, 6, 0, 2}

	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 10
5 6 U
4 7 U
8 3 L
8 3 L
1 10 U
9 2 U
10 1 L
10 1 L
8 3 U
8 3 U
`
	expect := []int{6, 7, 3, 0, 10, 2, 1, 0, 0, 0}

	runSample(t, s, expect)
}
