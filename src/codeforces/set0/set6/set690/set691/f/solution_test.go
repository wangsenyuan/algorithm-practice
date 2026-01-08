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
	s := `5
4 2 6 1 3
4
1 3 5 8
`
	expect := []int{20, 18, 14, 10}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
5 6
2
30 31
`
	expect := []int{2, 0}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2
2000000 2000000
1
2000000
`
	expect := []int{2}
	runSample(t, s, expect)
}
