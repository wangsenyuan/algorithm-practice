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
	s := `1
3 5 2 7
`
	expect := []int{5, 2, 7}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
2 2 9
3 3 1 4
`
	expect := []int{2, 9, 4}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
1 5
2 5 1
2 5 2
`
	expect := []int{5, 1}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3
3 4 4 9
7 7 6 5 4 3 2 1
4 2 4 5 1
`
	expect := []int{2, 4, 5, 1, 3, 2, 1}
	runSample(t, s, expect)
}
